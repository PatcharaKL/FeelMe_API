using dotnet.ViewModel;
using Project_FeelMe.Models;

namespace dotnet.Sevices.TokenService
{
    public interface ITokenService
    {
          Task<string> GeneraterTokenAccess(Account user);
          Task<AccountViewModels> DeCodeToken(string token);
          Task<string> GeneraterRefreshToken(Account user);
    }
}