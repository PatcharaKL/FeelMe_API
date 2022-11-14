
using dotnet.Sevices.TokenService;
using dotnet.ViewModel;
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
        private readonly FeelMeContext _dbContract;

        public UserController(ITokenService tokenService, IPassWordService passWordService, FeelMeContext dbContract)
        {
            _tokenService = tokenService;
            _passwordService = passWordService;
            _dbContract = dbContract;
        }
        [AllowAnonymous]
        [HttpPost("UserLogin")]
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
                    // var accessToken = await _tokenService.GeneraterTokenAccess(user);
                    // var refreshToken = await _tokenService.GeneraterRefreshToken();
                    return Ok(result);
                }
            }
            catch (Exception)
            {
                return Unauthorized("User not found");
            }

            return Unauthorized("User not found");
        }
        [HttpPost("UserLogOut")]
        public async Task<IActionResult> UserLogOut([FromBody] ResultToken token)
        {
            try
            {
                var user = await _tokenService.DeCodeToken(token.accessToken);
                var refreshTokenOut = await (from reToken in _dbContract.RefreshTokens
                                             where (reToken.AccountId == user.AccountId) && (reToken.IsValid == true)
                                             select new RefreshToken
                                             {
                                                 refreshToken = reToken.refreshToken,
                                                 AccountId = reToken.AccountId,
                                                 Exp = reToken.Exp,
                                                 IsValid = false
                                             }
                                           ).FirstOrDefaultAsync();
                _dbContract.Update(refreshTokenOut);
                await _dbContract.SaveChangesAsync();
                return Ok("Success");
            }
            catch (Exception)
            {
                return UnprocessableEntity();
            }

        }
        [HttpPost("NewTokenByRefreshToken")]
        public async Task<IActionResult> NewTokenByRefreshToken([FromBody] ResultToken resultToken)
        {
            try
            {
                var refreshTokenCk = await (from refreshToken in _dbContract.RefreshTokens
                                            where refreshToken.refreshToken == resultToken.refreshToken
                                            select new RefreshToken
                                            {
                                                refreshToken = refreshToken.refreshToken,
                                                AccountId = refreshToken.AccountId,
                                                Exp = refreshToken.Exp,
                                                IsValid = refreshToken.IsValid
                                            }).FirstOrDefaultAsync();
                if (refreshTokenCk.IsValid == true)
                {
                    RefreshToken refreshTokenUpdate = new RefreshToken
                    {
                        refreshToken = refreshTokenCk.refreshToken,
                        AccountId = refreshTokenCk.AccountId,
                        Exp = refreshTokenCk.Exp,
                        IsValid = false
                    };
                    _dbContract.Update(refreshTokenUpdate);
                    await _dbContract.SaveChangesAsync();
                }
                var userAccount = await (
                                          from account in _dbContract.Accounts
                                          where (account.AccountId == refreshTokenCk.AccountId)
                                          select new Account
                                          {
                                              AccountId = account.AccountId,
                                              Email = account.Email,
                                              PasswordHash = account.PasswordHash,
                                              Name = account.Name,
                                              Surname = account.Surname,
                                              Hp = account.Hp,
                                              Level = account.Level,
                                              PositionId = account.PositionId,
                                              DepartmentId = account.DepartmentId,
                                              CompanyId = account.CompanyId
                                          }).FirstOrDefaultAsync();
                var newResultToken = new ResultToken
                {
                    accessToken = await _tokenService.GeneraterTokenAccess(userAccount),
                    refreshToken = await _tokenService.GeneraterRefreshToken(userAccount)
                };

                return Ok(newResultToken);
            }
            catch(Exception)
            {
                return UnprocessableEntity();
            }
        }
        [Authorize]
        [HttpPost("[action]")]
        public async Task<IActionResult> GetUserDetail([FromBody] TokenSender.AccessToken token)
        {
            var data = await _tokenService.DeCodeToken(token.accessToken);
              var userAccount = await (
             from account in _dbContract.Accounts
             from position in _dbContract.Positions
             from depatrtment in _dbContract.Departments
             from company in _dbContract.Companies
             where (account.Email == data.Email)
                    &&(position.PositionId == data.PositionId)
                    &&(depatrtment.DepartmentId == data.DepartmentId)
                    &&(company.CompanyId == data.CompanyId)
             select new UserDetail
             {
                 Email = account.Email,
                 Name = account.Name,
                 Surname = account.Surname,
                 Hp = account.Hp,
                 Level = account.Level,
                 PositionName = position.PositionName,
                 DepartmentName = depatrtment.DepartmentName,
                 CompanyName = company.Name
             }).FirstOrDefaultAsync();
            return Ok(userAccount);
        }
      
        private async Task<Account> Authenticate(UserLogin userLogin)
        {
           
            var userAccount = await (
             from account in _dbContract.Accounts
             where (account.Email == userLogin.Email)
             select new Account
             {
                 AccountId = account.AccountId,
                 Email = account.Email,
                 PasswordHash = account.PasswordHash,
                 Name = account.Name,
                 Surname = account.Surname,
                 Hp = account.Hp,
                 Level = account.Level,
                 PositionId = account.PositionId,
                 DepartmentId = account.DepartmentId,
                 CompanyId = account.CompanyId
             }).FirstOrDefaultAsync();
              var ckPasswordHash = await _passwordService.CheckPassword(userLogin.Password,userAccount.PasswordHash);
            if (ckPasswordHash == true) return userAccount;
            else return null;
        }
    }
}