using System;
using System.Collections.Generic;

namespace FeelMe.Models
{
    public partial class Board
    {
        public int BoardId { get; set; }
        public string TitelBoard { get; set; } = null!;
        public string StoryBoard { get; set; } = null!;
        public DateTime Created { get; set; }
        public int AdcounId { get; set; }
        public int DepartmenId { get; set; }
    }
}
