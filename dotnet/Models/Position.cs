using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Position
    {
        public Position()
        {
            Accounts = new HashSet<Account>();
        }

        public int PositionId { get; set; }
        public string PositionName { get; set; }
        public DateTime Created { get; set; }

        public virtual ICollection<Account> Accounts { get; set; }
    }
}
