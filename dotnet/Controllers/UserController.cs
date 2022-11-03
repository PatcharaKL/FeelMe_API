
using dotnet.Sevices.TokenService;
using dotnet.ViewModel;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Project_FeelMe.Service.PassWordService;

namespace Project_FeelMe.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class UserController:ControllerBase
    {
         private readonly ITokenService _tokenService;
         private readonly IPassWordService _passwordService;

         public UserController(ITokenService tokenService,IPassWordService passWordService)
         {
            _tokenService = tokenService;
            _passwordService = passWordService;
         }
        [AllowAnonymous]
        [HttpPost("UserLogin")]
         public async Task<IActionResult> UserLogin([FromBody] UserLogin userLogin)
        {
                var user = Authenticate(userLogin);

            if (user != null)
            {
                var token = await _tokenService.GeneraterToken(user);
                return Ok(token);
            }

            return NotFound("User not found");
        }
        private  AccountViewModels  Authenticate (UserLogin userLogin)
        {

            var currentUser = UserConstants.Users.FirstOrDefault(o => o.Email.ToLower() == userLogin.Email.ToLower() && o.Password == userLogin.Password);

            if (currentUser != null)
            {
                return currentUser;
            }

            return null;
        }
    }
}