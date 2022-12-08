using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.LogDataService
{
    public interface ILogDataService
    {
         Task<List<Log>> GetAllLogDetailAsync();
         Task<List<Log>> GetAllLogDetailByAccountIdAsync(int accountId);
         Task SaveLogAsync(Log data);
    }
}