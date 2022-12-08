using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using dotnet.ViewModel;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Sevices.TokenService
{
    public class TokenService:ITokenService
    {
          private readonly IConfiguration _config;
           private readonly FeelMeContext _dbContract;
            public TokenService(IConfiguration config,FeelMeContext dbContract)
            {
                 _config = config;
                 _dbContract = dbContract;
            }
            public virtual async Task<string> GeneraterTokenAccess(Account user)
        {
            var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_config["Jwt:Key"]));
            var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);
            var claims = new[]
            {
                new Claim("Email", user.Email),
                new Claim("Name", user.Name),
                new Claim("Surname", user.Surname),
                new Claim("Role", user.PositionId.ToString()),
                new Claim("AccountId",user.AccountId.ToString()),
                new Claim("DepartmentId",user.DepartmentId.ToString()),
                new Claim("CompanyId",user.CompanyId.ToString())  
            };

            var token = new JwtSecurityToken(_config["Jwt:Issuer"],
              _config["Jwt:Audience"],
              claims,
              expires:DateTime.Now.AddSeconds(20),
              signingCredentials: credentials);


            return  await Task.FromResult<string>(new JwtSecurityTokenHandler().WriteToken(token));
        }
        public virtual async Task<AccountViewModels> DeCodeToken(string token)
        {
            var data = new JwtSecurityToken(token).Claims;
            var exp = int.Parse(data.FirstOrDefault(o => o.Type == "exp")?.Value);
            DateTimeOffset expires = DateTimeOffset.FromUnixTimeSeconds(exp).ToLocalTime();
            if(expires<DateTime.Now)
            {
                return null;
            }
            var user =  new AccountViewModels
                {
                   Name = data.FirstOrDefault(o => o.Type == "Name")?.Value,
                    Email = data.FirstOrDefault(o => o.Type == "Email")?.Value,
                    Surname = data.FirstOrDefault(o => o.Type == "Surname")?.Value,
                    PositionId = Convert.ToInt32(data.FirstOrDefault(o => o.Type == "Role")?.Value),
                    AccountId =  Convert.ToInt32(data.FirstOrDefault(o=>o.Type== "AccountId")?.Value),
                    DepartmentId = int.Parse(data.FirstOrDefault(o=>o.Type== "DepartmentId")?.Value),
                    CompanyId = int.Parse(data.FirstOrDefault(o=>o.Type== "CompanyId")?.Value)
                };
             return  await Task.FromResult(user);
        }
        public virtual async Task<string> GeneraterRefreshToken(Account user)
        {
            var refreshtoken = String.Join("-",Enumerable.Range(0,4).Select(options => Guid.NewGuid().ToString()).ToList()); 
                    RefreshToken refreshTokenAdd = new RefreshToken
                 {
                     refreshToken = refreshtoken,
                     AccountId = user.AccountId,
                     Exp = DateTime.Now.AddDays(15),
                    IsValid = true
                    };
                     _dbContract.Add(refreshTokenAdd);
                    await _dbContract.SaveChangesAsync();
          
            
            return await Task.FromResult<string>(refreshtoken);
        }
    }
}