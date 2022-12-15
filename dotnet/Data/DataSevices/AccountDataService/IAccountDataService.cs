using dotnet.ViewModel;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.AccountDataService
{
    public interface IAccountDataService
    {
         Task<Account> GetAccountByAccountIdAsync(int accountId);
         Task<UserDetail> GetUserDetailAsync(AccountViewModels data);
         Task<Account> GetAccountByEmailAsync(string email);
         Task UpdateAccountAsync(Account data);
         Task InsertAccountAsync(Account data);
         Task<List<UserDetail>> GetDetailEnemyAsync(int accountId);
         Task<List<UserDetail>> GetSearchAccountByNameAsync(string name);
    }
}