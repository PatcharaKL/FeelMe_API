using dotnet.ViewModel;


namespace dotnet.Sevices.TokenService
{
    public interface ITokenService
    {
          Task<string> GeneraterTokenAccess(AccountViewModels user);
          Task<AccountViewModels> DeCodeToken(string token);
          Task<string> GeneraterRefreshToken();
    }
}