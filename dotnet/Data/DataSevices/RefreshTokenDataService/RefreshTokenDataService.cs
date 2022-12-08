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
                var data = await( from refreshTokens in _dbContract.RefreshTokens
                                                         select new RefreshToken
                                                         {
                                                            refreshToken = refreshTokens.refreshToken,
                                                            AccountId = refreshTokens.AccountId,
                                                            Exp = refreshTokens.Exp,
                                                            IsValid = refreshTokens.IsValid
                                                         }).ToListAsync();
                return data;
            }
            public virtual async Task<RefreshToken> GetRefreshTokenByRefreshTokenAsync(string refreshToken)
            {
                  var data = await (from reToken in _dbContract.RefreshTokens
                                            where reToken.refreshToken == refreshToken
                                            select new RefreshToken
                                            {
                                                refreshToken = reToken.refreshToken,
                                                AccountId = reToken.AccountId,
                                                Exp = reToken.Exp,
                                                IsValid = reToken.IsValid
                                            }).FirstOrDefaultAsync();
                 return data;                      
            }
             public virtual async Task UpdateRefreshTokenAsync(RefreshToken refreshToken)
             {
                            _dbContract.Update(refreshToken);
                            await _dbContract.SaveChangesAsync();
             }  
    }
}