using System;
using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata;
using Project_FeelMe.Models;

namespace Project_FeelMe.Data
{
    public partial class FeelMeContext : DbContext
    {
        public FeelMeContext()
        {
        }

        public FeelMeContext(DbContextOptions<FeelMeContext> options)
            : base(options)
        {
        }

        public virtual DbSet<Account> Accounts { get; set; }
        public virtual DbSet<Board> Boards { get; set; }
        public virtual DbSet<Comment> Comments { get; set; }
        public virtual DbSet<Company> Companies { get; set; }
        public virtual DbSet<Department> Departments { get; set; }
        public virtual DbSet<Log> Logs { get; set; }
        public virtual DbSet<Position> Positions { get; set; }
        public virtual DbSet<RefreshToken> RefreshTokens { get; set; }
        public virtual DbSet<Weapon> Weapons { get; set; }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            if (!optionsBuilder.IsConfigured)
            {
                optionsBuilder.UseMySql("name=DefaultConnectionString", Microsoft.EntityFrameworkCore.ServerVersion.Parse("10.6.9-mariadb"));
            }
        }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.UseCollation("utf8mb3_general_ci")
                .HasCharSet("utf8mb3");

            modelBuilder.Entity<Account>(entity =>
            {
                entity.HasKey(e => new { e.AccountId, e.DepartmentId, e.PositionId, e.CompanyId })
                    .HasName("PRIMARY")
                    .HasAnnotation("MySql:IndexPrefixLength", new[] { 0, 0, 0, 0 });

                entity.ToTable("accounts");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.HasIndex(e => e.PositionId, "company_id_idx");

                entity.HasIndex(e => e.CompanyId, "company_id_idx1");

                entity.HasIndex(e => e.DepartmentId, "department_id_idx");

                entity.Property(e => e.AccountId)
                    .HasColumnType("int(11)")
                    .ValueGeneratedOnAdd()
                    .HasColumnName("account_id");

                entity.Property(e => e.DepartmentId)
                    .HasColumnType("int(11)")
                    .HasColumnName("department_id");

                entity.Property(e => e.PositionId)
                    .HasColumnType("int(11)")
                    .HasColumnName("position_id");

                entity.Property(e => e.CompanyId)
                    .HasColumnType("int(11)")
                    .HasColumnName("company_id");

                entity.Property(e => e.ApplyDate)
                    .HasColumnType("datetime")
                    .HasColumnName("apply_date");

                entity.Property(e => e.AvatarUrl)
                    .HasMaxLength(100)
                    .HasColumnName("avatar_url");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.Email)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("email");

                entity.Property(e => e.Hp)
                    .HasColumnType("int(11)")
                    .HasColumnName("hp");

                entity.Property(e => e.IsActive)
                    .IsRequired()
                    .HasColumnName("is_active")
                    .HasDefaultValueSql("'1'");

                entity.Property(e => e.Level)
                    .HasColumnType("int(11)")
                    .HasColumnName("level");

                entity.Property(e => e.Name)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("name");

                entity.Property(e => e.PasswordHash)
                    .IsRequired()
                    .HasMaxLength(256)
                    .HasColumnName("password_hash");

                entity.Property(e => e.Surname)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("surname");

                entity.HasOne(d => d.Company)
                    .WithMany(p => p.Accounts)
                    .HasForeignKey(d => d.CompanyId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("compony_id");

                entity.HasOne(d => d.Position)
                    .WithMany(p => p.Accounts)
                    .HasForeignKey(d => d.PositionId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("position_id");
            });

            modelBuilder.Entity<Board>(entity =>
            {
                entity.HasKey(e => new { e.BoardId, e.AdcounId, e.DepartmenId })
                    .HasName("PRIMARY")
                    .HasAnnotation("MySql:IndexPrefixLength", new[] { 0, 0, 0 });

                entity.ToTable("boards");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.HasIndex(e => e.AdcounId, "adcount_id_idx");

                entity.HasIndex(e => e.DepartmenId, "department_id_idx");

                entity.Property(e => e.BoardId)
                    .HasColumnType("int(11)")
                    .HasColumnName("board_id");

                entity.Property(e => e.AdcounId)
                    .HasColumnType("int(11)")
                    .HasColumnName("adcoun_id");

                entity.Property(e => e.DepartmenId)
                    .HasColumnType("int(11)")
                    .HasColumnName("departmen_id");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.StoryBoard)
                    .IsRequired()
                    .HasColumnType("mediumtext")
                    .HasColumnName("story_board");

                entity.Property(e => e.TitelBoard)
                    .IsRequired()
                    .HasMaxLength(500)
                    .HasColumnName("titel_board");
            });

            modelBuilder.Entity<Comment>(entity =>
            {
                entity.HasKey(e => new { e.CommentId, e.AccountId, e.BoradId })
                    .HasName("PRIMARY")
                    .HasAnnotation("MySql:IndexPrefixLength", new[] { 0, 0, 0 });

                entity.ToTable("comments");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.HasIndex(e => e.AccountId, "account_fk_idx");

                entity.HasIndex(e => e.BoradId, "borad_fk_idx");

                entity.Property(e => e.CommentId)
                    .HasColumnType("int(11)")
                    .HasColumnName("comment_id");

                entity.Property(e => e.AccountId)
                    .HasColumnType("int(11)")
                    .HasColumnName("account_id");

                entity.Property(e => e.BoradId)
                    .HasColumnType("int(11)")
                    .HasColumnName("borad_id");

                entity.Property(e => e.CommentText)
                    .IsRequired()
                    .HasColumnType("mediumtext")
                    .HasColumnName("comment_text");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.EditTime)
                    .HasColumnType("datetime")
                    .HasColumnName("edit_time");
            });

            modelBuilder.Entity<Company>(entity =>
            {
                entity.ToTable("companies");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.Property(e => e.CompanyId)
                    .HasColumnType("int(11)")
                    .HasColumnName("company_id");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.CreatorId)
                    .HasMaxLength(45)
                    .HasColumnName("creator_id");

                entity.Property(e => e.Name)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("name");

                entity.Property(e => e.RoomName)
                    .HasMaxLength(100)
                    .HasColumnName("room_name");
            });

            modelBuilder.Entity<Department>(entity =>
            {
                entity.HasKey(e => new { e.DepartmentId, e.CompanyId })
                    .HasName("PRIMARY")
                    .HasAnnotation("MySql:IndexPrefixLength", new[] { 0, 0 });

                entity.ToTable("departments");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.HasIndex(e => e.CompanyId, "company_id_idx");

                entity.Property(e => e.DepartmentId)
                    .HasColumnType("int(11)")
                    .HasColumnName("department_id");

                entity.Property(e => e.CompanyId)
                    .HasColumnType("int(11)")
                    .HasColumnName("company_id");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.DepartmentName)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("department_name");

                entity.HasOne(d => d.Company)
                    .WithMany(p => p.Departments)
                    .HasForeignKey(d => d.CompanyId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("company_id");
            });

            modelBuilder.Entity<Log>(entity =>
            {
                entity.HasKey(e => new { e.LogId, e.AccountId })
                    .HasName("PRIMARY")
                    .HasAnnotation("MySql:IndexPrefixLength", new[] { 0, 0 });

                entity.ToTable("logs");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.HasIndex(e => e.AccountId, "account_id_idx");

                entity.Property(e => e.LogId)
                    .HasColumnType("int(11)")
                    .ValueGeneratedOnAdd()
                    .HasColumnName("log_id");

                entity.Property(e => e.AccountId)
                    .HasColumnType("int(11)")
                    .HasColumnName("account_id");

                entity.Property(e => e.Amount)
                    .HasColumnType("int(11)")
                    .HasColumnName("amount");

                entity.Property(e => e.Datetime)
                    .HasColumnType("datetime")
                    .HasColumnName("datetime")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.Type)
                    .HasColumnType("tinyint(4)")
                    .HasColumnName("type");
            });

            modelBuilder.Entity<Position>(entity =>
            {
                entity.ToTable("positions");

                entity.HasCharSet("utf8mb4")
                    .UseCollation("utf8mb4_unicode_ci");

                entity.Property(e => e.PositionId)
                    .HasColumnType("int(11)")
                    .HasColumnName("position_id");

                entity.Property(e => e.Created)
                    .HasColumnType("datetime")
                    .HasColumnName("created")
                    .HasDefaultValueSql("current_timestamp()");

                entity.Property(e => e.PositionName)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("position_name");
            });

            modelBuilder.Entity<RefreshToken>(entity =>
            {
                entity.HasKey(e => e.refreshToken)
                    .HasName("PRIMARY");

                entity.ToTable("refresh_token");

                entity.HasIndex(e => e.AccountId, "account_id");

                entity.Property(e => e.refreshToken).HasColumnName("refreshToken");

                entity.Property(e => e.AccountId)
                    .HasColumnType("int(11)")
                    .HasColumnName("account_id");

                entity.Property(e => e.Exp)
                    .HasColumnType("datetime")
                    .HasColumnName("exp");

                entity.Property(e => e.IsValid)
                    .IsRequired()
                    .HasColumnName("isValid")
                    .HasDefaultValueSql("'1'");
            });

            modelBuilder.Entity<Weapon>(entity =>
            {
                entity.HasKey(e => e.WeaponsId)
                    .HasName("PRIMARY");

                entity.ToTable("weapons");

                entity.Property(e => e.WeaponsId)
                    .HasColumnType("int(11)")
                    .HasColumnName("weapons_id");

                entity.Property(e => e.UrlWeapon)
                    .IsRequired()
                    .HasMaxLength(200);

                entity.Property(e => e.WeaponName)
                    .IsRequired()
                    .HasMaxLength(100)
                    .HasColumnName("weaponName");
            });

            OnModelCreatingPartial(modelBuilder);
        }

        partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
    }
}
