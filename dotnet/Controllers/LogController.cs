using dotnet.Data.DataSevices.AccountDataService;
using dotnet.Data.DataSevices.LogDataService;
using dotnet.Sevices.TokenService;
using dotnet.ViewModel;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Controllers
{
    [Authorize]
    [ApiController]
    [Route("[controller]")]
    public class LogController : ControllerBase
    {
          
          private readonly ILogDataService _logDataService;
          private readonly IAccountDataService _accountDataService;
          
          private readonly ITokenService _tokenService;

          public LogController(ILogDataService logDataService,IAccountDataService accountDataService,ITokenService tokenService)
          {
            _logDataService = logDataService;
            _accountDataService = accountDataService;
            _tokenService = tokenService;
          }
    
          [HttpPost("[action]")]
          public async Task<IActionResult> AttackDamage([FromBody] LogSender  logSender)
          {
            try
            {
                 var token  = HttpContext.GetTokenAsync("access_token").Result;
                 var data = await _tokenService.DeCodeToken(token);
                 Log dataLog = new Log
                 {
                    AccountId = data.AccountId,
                    Datetime = DateTime.Now,
                    Type = logSender.type,
                    Amount = logSender.amount
                 };
                 var userUpdate = await _accountDataService.GetAccountByAccountIdAsync(dataLog.AccountId);
                 await _logDataService.SaveLogAsync(dataLog);
                 userUpdate.Hp -= dataLog.Amount;
                 await _accountDataService.UpdateAsync(userUpdate);
                 return Ok("Success");
            }
            catch(Exception e)
            {
                return UnprocessableEntity(e.Message);
            }
           
          }

    }
}