using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using dotnet.ViewModel;
using Microsoft.IdentityModel.Tokens;


namespace dotnet.Sevices.TokenService
{
    public class TokenService:ITokenService
    {
          private IConfiguration _config;
            public TokenService(IConfiguration config)
            {
                 _config = config;
            }
            public virtual async Task<string> GeneraterTokenAccess(AccountViewModels user)
        {
            var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_config["Jwt:Key"]));
            var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);

            var claims = new[]
            {
                new Claim(ClaimTypes.NameIdentifier, user.Email),
                new Claim(ClaimTypes.Email, user.Email),
                new Claim(ClaimTypes.Name, user.Name),
                new Claim(ClaimTypes.Surname, user.Surname),
                new Claim(ClaimTypes.Role, user.PositionId.ToString())
               
            };

            var token = new JwtSecurityToken(_config["Jwt:Issuer"],
              _config["Jwt:Audience"],
              claims,
              signingCredentials: credentials);

            return  await Task.FromResult<string>(new JwtSecurityTokenHandler().WriteToken(token));
        }
        public virtual async Task<AccountViewModels> DeCodeToken(string token)
        {
            var data = new JwtSecurityToken(token).Claims;
            var user =  new AccountViewModels
                {
                   Name = data.FirstOrDefault(o => o.Type == ClaimTypes.Name)?.Value,
                    Email = data.FirstOrDefault(o => o.Type == ClaimTypes.NameIdentifier)?.Value,
                    Surname = data.FirstOrDefault(o => o.Type == ClaimTypes.Surname)?.Value,
                    PositionId = Convert.ToInt32(data.FirstOrDefault(o => o.Type == ClaimTypes.Role)?.Value)
                };
             return  await Task.FromResult(user);
        }
        public virtual async Task<string> GeneraterRefreshToken()
        {
            return await Task.FromResult<string>(String.Join("-",Enumerable.Range(0,4).Select(options => Guid.NewGuid().ToString()).ToList()));
        }
    }
}