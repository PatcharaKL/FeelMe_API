using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.RefreshTokenDataService
{
    public class RefreshTokenDataService:IRefreshTokenDataService
    {
            private readonly FeelMeContext _dbContract;

            public RefreshTokenDataService(FeelMeContext dbContract)
            {
                 _dbContract = dbContract;
            }

            public virtual async Task<List<RefreshToken>> GetAllRefreshTokenAsync()
            {
                var data = await _dbContract.RefreshTokens.ToListAsync();
                return data;
            }
            public  virtual async Task<RefreshToken> GetRefreshTokenListByAccountIdAsync(int accountId )
            {
                return await _dbContract.RefreshTokens.FirstOrDefaultAsync(reToken => reToken.AccountId == accountId && reToken.IsValid == true);
            }
            public virtual async Task<RefreshToken> GetRefreshTokenByRefreshTokenAsync(string refreshToken)
            {
                  var data = await _dbContract.RefreshTokens.FirstOrDefaultAsync(reToken => reToken.refreshToken ==refreshToken);
                 return data;                      
            }
             public virtual async Task UpdateRefreshTokenAsync(RefreshToken refreshToken)
             {
                 _dbContract.Update(refreshToken);
                await _dbContract.SaveChangesAsync();
             }
             public virtual async Task UpdateListRefreshTokenAsync(List<RefreshToken> refreshToken)
             {
                 _dbContract.UpdateRange(refreshToken);
                await _dbContract.SaveChangesAsync();
             }   
             public virtual async Task InsertAsyncRefreshToken(RefreshToken refreshToken)
             {
                _dbContract.Add(refreshToken);
                await _dbContract.SaveChangesAsync();
             }
    }
}