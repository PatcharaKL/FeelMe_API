namespace dotnet.ViewModel
{
    public class TokenSender
    {
        public TokenSender()
        {
            accessTokens = new List<AccessToken>();
            refreshTokens = new List<RefreshToken>();
        }
        public string tokenSend { get; set; }
        public List<AccessToken> accessTokens {get;set;}
         public List<RefreshToken> refreshTokens {get;set;}
        public class AccessToken
        {
             public string accessToken{get;set;}
        }
         public class RefreshToken
        {
             public string refreshToken{get;set;}
        }
    }
}