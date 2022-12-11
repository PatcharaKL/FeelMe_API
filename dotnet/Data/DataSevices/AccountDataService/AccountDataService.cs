using dotnet.ViewModel;
using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.AccountDataService
{
    public class AccountDataService:IAccountDataService
    {
          private readonly FeelMeContext _dbContract;

            public AccountDataService(FeelMeContext dbContract)
            {
                 _dbContract = dbContract;
            }
             public virtual async Task<Account> GetAccountByAccountIdAsync(int accountId)
             {
                 var data =  await _dbContract.Accounts.FirstOrDefaultAsync(account => account.AccountId == accountId);
                return data;
             }
              public virtual async Task<UserDetail> GetUserDetailAsync(AccountViewModels data)
              {
                    
                     var newData = await (
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
                    return newData;
              }
              public virtual async Task<List<UserDetail>> GetDetailEnemyAsync(int accountId)
              {
                     var newData = await (
                                                 from account in _dbContract.Accounts
                                                 from position in _dbContract.Positions
                                                 from depatrtment in _dbContract.Departments
                                                 from company in _dbContract.Companies
                                                 where (account.AccountId != accountId) 
                                                     &&(position.PositionId == account.PositionId)
                                                     &&(depatrtment.DepartmentId == account.DepartmentId)
                                                     &&(company.CompanyId == account.CompanyId)
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
                                             }).ToListAsync();
                    return newData;
              }
               public virtual async Task<Account> GetAccountByEmailAsync(string email)
               {
                     var data = await _dbContract.Accounts.FirstOrDefaultAsync(account => account.Email == email);
                    return data;
               }
               public virtual async Task UpdateAccountAsync(Account data)
               {
                      _dbContract.Update(data);
                      await _dbContract.SaveChangesAsync();
               }
                public virtual async Task InsertAccountAsync(Account data)
             {
                _dbContract.Add(data);
                await _dbContract.SaveChangesAsync();
             }
    }
}