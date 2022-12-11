namespace dotnet.ViewModel
{
    public class ResultToken
    {
        public ResultToken()
        {
            refreshTokens =  new List<RefreshToken>();
            accessTokens = new List<AccessToken>();
        }
        public string accessToken{get;set;}
        public string refreshToken{get;set;}
        public List<RefreshToken> refreshTokens{get;set;}
        public List<AccessToken> accessTokens{get;set;}
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