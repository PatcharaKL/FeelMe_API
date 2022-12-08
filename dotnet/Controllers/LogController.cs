using dotnet.Data.DataSevices.AccountDataService;
using dotnet.Data.DataSevices.LogDataService;
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

          public LogController(ILogDataService logDataService,IAccountDataService accountDataService)
          {
            _logDataService = logDataService;
            _accountDataService = accountDataService;
          }
    
          [HttpPost("[action]")]
          public async Task<IActionResult> AttackDamage([FromBody] Log  dataLog)
          {
            try
            {
                 var userUpdate = await _accountDataService.GetAccountByAccountIdAsync(dataLog.AccountId);
                 await _logDataService.SaveLogAsync(dataLog);
                 userUpdate.Hp -= dataLog.Amount;
                 await _accountDataService.UpdateAsync(userUpdate);
                 return Ok("Success");
            }
            catch(Exception e)
            {
                return Unauthorized(e.Message);
            }
           
          }

    }
}