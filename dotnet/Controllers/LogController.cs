using dotnet.Data.DataSevices.AccountDataService;
using dotnet.Data.DataSevices.LogDataService;
using dotnet.Sevices.TokenService;
using dotnet.ViewModel;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Project_FeelMe.Data;
using Project_FeelMe.Models;
using dotnet.Data.DataSevices.WeaponDataService;

namespace dotnet.Controllers
{
   
    [ApiController]
    [Route("[controller]")]
    public class LogController : ControllerBase
    {
          
          private readonly ILogDataService _logDataService;
          private readonly IAccountDataService _accountDataService;
          private readonly IWeaponDataService _weaponDataService;
          private readonly ITokenService _tokenService;

          public LogController(ILogDataService logDataService,IAccountDataService accountDataService,ITokenService tokenService,IWeaponDataService weaponDataService)
          {
            _logDataService = logDataService;
            _accountDataService = accountDataService;
            _tokenService = tokenService;
            _weaponDataService = weaponDataService;
          }
    
          [HttpPost("[action]")]
          [Authorize]
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
                 if(userUpdate.Hp-logSender.amount>=0)
                 {
                     await _logDataService.SaveLogAsync(dataLog);
                     userUpdate.Hp -= dataLog.Amount;
                     await _accountDataService.UpdateAccountAsync(userUpdate);
                     return Ok();
                 }
                 else
                 {
                     await _logDataService.SaveLogAsync(dataLog);
                     userUpdate.Hp = 0;
                     await _accountDataService.UpdateAccountAsync(userUpdate);
                     return  Ok();
                 }
                 
             }
             catch(Exception)
             {
                return Unauthorized();
             }
          }
          [HttpPost("[action]")]
          [Authorize]
           public async Task<IActionResult> HealingUpSender([FromBody] LogSender  logSender)
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
                 await _logDataService.SaveLogAsync(dataLog);
                return Ok();
             }
             catch(Exception)
             {
               return Unauthorized();
             }
           }
          [HttpPost("[action]")]
          [Authorize]
           public async Task<IActionResult> HealingUp([FromBody] LogSender  logSender)
           {
             try
             {
                var token  = HttpContext.GetTokenAsync("access_token").Result;
                var data = await _tokenService.DeCodeToken(token);
                 
                return Ok();
             }
             catch(Exception)
             {
               return Unauthorized();
             }
           }
         [HttpPost("[action]")]
           public async Task<IActionResult> GetWeapons()
           {
               var data = await _weaponDataService.GetAllWeaponAsync();
               return Ok(data);
           }
    }
}