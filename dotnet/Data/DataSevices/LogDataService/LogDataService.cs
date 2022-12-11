using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.LogDataService
{
    public class LogDataService:ILogDataService
    {
            private readonly FeelMeContext _dbContract;

              public LogDataService(FeelMeContext dbContract)
            {
                 _dbContract = dbContract;
            }
             public virtual async Task<List<Log>> GetAllLogDetailAsync()
            {
                var data = await _dbContract.Logs.ToListAsync();
                return data;
            }
             public virtual async Task<List<Log>> GetAllLogDetailByAccountIdAsync(int accountId)
            {
                var data = await( from logs in _dbContract.Logs
                                                         where logs.AccountId == accountId
                                                         select new Log
                                                         {
                                                           LogId = logs.LogId,
                                                           Type = logs.Type,
                                                           Amount = logs.Amount,
                                                           Datetime = logs.Datetime,
                                                           AccountId = logs.AccountId
                                                         }).ToListAsync();
                return data;
            }
            public virtual async Task SaveLogAsync(Log data)
            {
                _dbContract.Add(data);
                await _dbContract.SaveChangesAsync();
            }

    }
}