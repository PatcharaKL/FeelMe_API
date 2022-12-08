using System.Text;
using dotnet.Data.DataSevices.AccountDataService;
using dotnet.Data.DataSevices.LogDataService;
using dotnet.Data.DataSevices.RefreshTokenDataService;
using dotnet.Sevices.TokenService;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using Project_FeelMe.Data;
using Project_FeelMe.Service.PassWordService;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddHealthChecks();
builder.Services.AddSwaggerGen();
builder.Services.AddTransient<IPassWordService,PassWordService>();
builder.Services.AddTransient<ITokenService,TokenService>();
builder.Services.AddTransient<IRefreshTokenDataService,RefreshTokenDataService>();
builder.Services.AddTransient<IAccountDataService,AccountDataService>();
builder.Services.AddTransient<ILogDataService,LogDataService>();
 builder.Services.AddDbContext<FeelMeContext>(options => options.UseMySQL(builder.Configuration.GetConnectionString("DefaultConnectionString")));
 builder.Services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme)
                .AddJwtBearer(options => {
                    options.TokenValidationParameters = new TokenValidationParameters
                    {
                        ValidateIssuer = true,
                        ValidateAudience = true,
                        ValidateLifetime = true,
                        ValidateIssuerSigningKey = true,
                        ValidIssuer =builder.Configuration["Jwt:Issuer"],
                        ValidAudience = builder.Configuration["Jwt:Audience"],
                        IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(builder.Configuration["Jwt:Key"]))
                    };
                });

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseAuthentication();

app.UseAuthorization();

app.MapHealthChecks("/healthCheck");

app.MapControllers();

app.Run();
