using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Company
    {
        public Company()
        {
            Accounts = new HashSet<Account>();
            Departments = new HashSet<Department>();
        }

        public int CompanyId { get; set; }
        public string Name { get; set; }
        public string RoomName { get; set; }
        public string CreatorId { get; set; }
        public DateTime Created { get; set; }

        public virtual ICollection<Account> Accounts { get; set; }
        public virtual ICollection<Department> Departments { get; set; }
    }
}
