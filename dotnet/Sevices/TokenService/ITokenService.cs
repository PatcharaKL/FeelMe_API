using dotnet.ViewModel;


namespace dotnet.Sevices.TokenService
{
    public interface ITokenService
    {
          Task<string> GeneraterToken(AccountViewModels user);
          Task<AccountViewModels> DeCodeToken(string token);
    }
}