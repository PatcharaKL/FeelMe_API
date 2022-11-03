using System;
using System.Collections.Generic;

namespace FeelMe.Models
{
    public partial class Account
    {
        public int AccountId { get; set; }
        public string Email { get; set; } = null!;
        public string PasswordHash { get; set; } = null!;
        public string Name { get; set; } = null!;
        public string Surname { get; set; } = null!;
        public string AvatarUrl { get; set; } = null!;
        public DateTime ApplyDate { get; set; }
        public string IsActive { get; set; } = null!;
        public int Hp { get; set; }
        public int Level { get; set; }
        public DateTime Created { get; set; }
        public int DepartmentId { get; set; }
        public int PositionId { get; set; }
        public int CompanyId { get; set; }

        public virtual Company Company { get; set; } = null!;
        public virtual Position Position { get; set; } = null!;
    }
}
