using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Board
    {
        public int BoardId { get; set; }
        public string TitelBoard { get; set; }
        public string StoryBoard { get; set; }
        public DateTime Created { get; set; }
        public int AdcounId { get; set; }
        public int DepartmenId { get; set; }
    }
}
