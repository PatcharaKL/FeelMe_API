using System;
using System.Collections.Generic;

namespace FeelMe.Models
{
    public partial class Log
    {
        public int LogId { get; set; }
        public sbyte Type { get; set; }
        public int Amount { get; set; }
        public DateTime Datetime { get; set; }
        public int AccountId { get; set; }
    }
}
