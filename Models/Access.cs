using System;
using System.Collections.Generic;

namespace FeelMe.Models
{
    public partial class Access
    {
        public int AccesseId { get; set; }
        public DateTime Created { get; set; }
        public int AdcountId { get; set; }
    }
}
