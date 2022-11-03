using System;
using System.Collections.Generic;

namespace FeelMe.Models
{
    public partial class Company
    {
        public Company()
        {
            Accounts = new HashSet<Account>();
            Departments = new HashSet<Department>();
        }

        public int CompanyId { get; set; }
        public string Surname { get; set; } = null!;
        public string RoomName { get; set; } = null!;
        public string CreatorId { get; set; } = null!;
        public DateTime Created { get; set; }

        public virtual ICollection<Account> Accounts { get; set; }
        public virtual ICollection<Department> Departments { get; set; }
    }
}
