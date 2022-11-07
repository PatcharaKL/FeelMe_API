using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Account
    {
        public int AccountId { get; set; }
        public string Email { get; set; }
        public string PasswordHash { get; set; }
        public string Name { get; set; }
        public string Surname { get; set; }
        public string AvatarUrl { get; set; }
        public DateTime ApplyDate { get; set; }
        public bool? IsActive { get; set; }
        public int Hp { get; set; }
        public int Level { get; set; }
        public DateTime Created { get; set; }
        public int DepartmentId { get; set; }
        public int PositionId { get; set; }
        public int CompanyId { get; set; }

        public virtual Company Company { get; set; }
        public virtual Position Position { get; set; }
    }
}
