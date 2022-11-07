using System;
using System.Collections.Generic;

namespace Project_FeelMe.Models
{
    public partial class Comment
    {
        public int CommentId { get; set; }
        public string CommentText { get; set; }
        public DateTime Created { get; set; }
        public DateTime? EditTime { get; set; }
        public int AccountId { get; set; }
        public int BoradId { get; set; }
    }
}
