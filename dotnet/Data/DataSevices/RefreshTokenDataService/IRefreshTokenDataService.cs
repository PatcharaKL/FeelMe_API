using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.RefreshTokenDataService
{
    public interface IRefreshTokenDataService
    {
         Task<List<RefreshToken>> GetAllRefreshTokenAsync();
         Task<RefreshToken> GetRefreshTokenByRefreshTokenAsync(string refreshToken);
         Task UpdateRefreshTokenAsync(RefreshToken refreshToken);
         Task InsertAsyncRefreshToken(RefreshToken refreshToken);
        Task<RefreshToken> GetRefreshTokenListByAccountIdAsync(int accountId );
         Task UpdateListRefreshTokenAsync(List<RefreshToken> refreshToken);
    }
}