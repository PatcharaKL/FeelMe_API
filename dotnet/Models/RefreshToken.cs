using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class RefreshToken
    {
        public string refreshToken { get; set; }
        public int AccountId { get; set; }
        public DateTime Exp { get; set; }
        public bool? IsValid { get; set; }
    }
}
