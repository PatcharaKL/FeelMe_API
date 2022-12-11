namespace dotnet.ViewModel
{
    public class ResultToken
    {
        public ResultToken()
        {
            refreshTokens =  new RefreshToken();
            accessTokens = new AccessToken();
            resulToken = new TokenSender();
        }
     
        public RefreshToken refreshTokens{get;set;}
        public AccessToken accessTokens{get;set;}
        public TokenSender resulToken{get;set;}
        public class TokenSender
        {
               public string accessToken{get;set;}
               public string refreshToken{get;set;}
        }
        public class RefreshToken
        {
            public string refreshToken{get;set;}
        }
         public class AccessToken
        {
            public string accessToken{get;set;}
        }
    }
}