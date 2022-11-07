using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Department
    {
        public int DepartmentId { get; set; }
        public string DepartmentName { get; set; }
        public DateTime Created { get; set; }
        public int CompanyId { get; set; }

        public virtual Company Company { get; set; }
    }
}
