
using System.Security.Claims;
using dotnet.Data.DataSevices.AccountDataService;
using dotnet.Data.DataSevices.RefreshTokenDataService;
using dotnet.Sevices.TokenService;
using dotnet.ViewModel;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;
using Project_FeelMe.Service.PassWordService;

namespace Project_FeelMe.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class UserController : ControllerBase
    {
        private readonly ITokenService _tokenService;
        private readonly IPassWordService _passwordService;
        private readonly IRefreshTokenDataService _refreshTokenDataService;
        private readonly IAccountDataService _accountDataService;

        public UserController
        (
            ITokenService tokenService, IPassWordService passWordService,
            IRefreshTokenDataService refreshTokenDataService,
            IAccountDataService accountDataService
        )
        {
            _tokenService = tokenService;
            _passwordService = passWordService;
            _refreshTokenDataService = refreshTokenDataService;
            _accountDataService = accountDataService;
        }
        [AllowAnonymous]
        [HttpPost("[action]")]
        public async Task<IActionResult> UserLogin([FromBody] UserLogin userLogin)
        {
            try
            {
                var user = await Authenticate(userLogin);

                if (user != null)
                {
                    var result = new ResultToken
                    {
                        accessToken = await _tokenService.GeneraterTokenAccess(user),
                        refreshToken = await _tokenService.GeneraterRefreshToken(user)
                    };

                    return Ok(result);
                }
            }
            catch (Exception)
            {
                return Unauthorized("User not found");
            }

            return Unauthorized("User not found");
        }
        [Authorize]
        [HttpPost("[action]")]
        public async Task<IActionResult> UserLogOut([FromBody] ResultToken.RefreshToken resultToken)
        {
            try
            {
                var token  = HttpContext.GetTokenAsync("access_token").Result;
                var user = await _tokenService.DeCodeToken(token);
                var reToken = await _refreshTokenDataService.GetRefreshTokenByRefreshTokenAsync(resultToken.refreshToken);
                reToken.IsValid = false;
                await  _refreshTokenDataService.UpdateRefreshTokenAsync(reToken);
                return Ok("Success");
            }
            catch (Exception)
            {
                return UnprocessableEntity();
            }

        }
        [HttpPost("[action]")]
        public async Task<IActionResult> NewTokenByRefreshToken([FromBody] ResultToken.RefreshToken resultToken)
        {
          
            var refreshTokenCk = await _refreshTokenDataService.GetRefreshTokenByRefreshTokenAsync(resultToken.refreshToken);
            if(refreshTokenCk.Exp < DateTime.Now)
            {
                refreshTokenCk.IsValid = false;                
                await _refreshTokenDataService.UpdateRefreshTokenAsync(refreshTokenCk);
                 return  Unauthorized("ReToken is Exp");
            }
           else if (refreshTokenCk.IsValid == true )
            {
                var refTokenLists = await _refreshTokenDataService.GetRefreshTokenListByAccountIdAsync(refreshTokenCk.AccountId);
                foreach(RefreshToken re in refTokenLists)
                {
                       re.IsValid = false;
                       await _refreshTokenDataService.UpdateRefreshTokenAsync(re);
                }
                var userAccount = await _accountDataService.GetAccountByAccountIdAsync(refreshTokenCk.AccountId);
                var newResultToken = new ResultToken
                {
                    accessToken = await _tokenService.GeneraterTokenAccess(userAccount),
                    refreshToken = await _tokenService.GeneraterRefreshToken(userAccount)
                };

                return Ok(newResultToken);
            }
            else return Unauthorized();
        }

        [HttpPost("[action]")]
        public async Task<IActionResult> GetRefreshTokenDetail()
        {
            var data =  await _refreshTokenDataService.GetRefreshTokenListByAccountIdAsync(1); 
            return Ok(data);
        }
        [Authorize]
        [HttpPost("[action]")]
        public async Task<IActionResult> GetUserDetail()
        {
            var token  = HttpContext.GetTokenAsync("access_token").Result;
            var data = await _tokenService.DeCodeToken(token);
            if(data==null) return Unauthorized();
              var userAccount = await _accountDataService.GetUserDetailAsync(data);
            return Ok(userAccount);
        }
      
        private async Task<Account> Authenticate(UserLogin userLogin)
        {
            var userAccount = await _accountDataService.GetAccountByEmailAsync(userLogin.Email);
              var ckPasswordHash = await _passwordService.CheckPassword(userLogin.Password,userAccount.PasswordHash);
            if (ckPasswordHash == true) return userAccount;
            else return null;
        }
    }
}