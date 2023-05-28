package main

import (
	"giasuaeapi/src/config"
	"giasuaeapi/src/controllers"
	"giasuaeapi/src/middleware"
	"giasuaeapi/src/repositories"
	"giasuaeapi/src/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	accountReponsitory   repositories.AccountReponsitory   = repositories.NewAccountReponsitory(db)
	sugbjectRepository   repositories.SubjectRepository    = repositories.NewSubjectRepository(db)
	newClassRepository   repositories.NewClassRepository   = repositories.NewNewClassRepository(db)
	classRepository      repositories.ClassRepository      = repositories.NewClassITRepository(db)
	categoryRepository   repositories.CategoryRepository   = repositories.NewCategoryRepository(db)
	postRepository       repositories.PostRepository       = repositories.NewPostRepository(db)
	transRepository      repositories.TransRepository      = repositories.NewTransRepository(db)
	salaryinfoRepository repositories.SalaryinfoRepository = repositories.NewSalaryinfoRepository(db)
	tutorRepository      repositories.TutorRepository      = repositories.NewTutorRepository(db)

	jwtService        services.JWTService        = services.NewJWTService()
	accountService    services.AccountService    = services.NewAccountService(accountReponsitory)
	subjectService    services.SubjectService    = services.NewSubjectService(sugbjectRepository)
	authService       services.AuthService       = services.NewAuthService(accountReponsitory)
	newClassService   services.NewClassService   = services.NewNewClassService(newClassRepository)
	classService      services.ClassService      = services.NewClassITService(classRepository)
	categoryService   services.CategoryService   = services.NewCategoryService(categoryRepository)
	postService       services.PostService       = services.NewPostService(postRepository)
	transService      services.TransService      = services.NewTransService(transRepository)
	salaryinfoService services.SalaryinfoService = services.NewSalaryinfoService(salaryinfoRepository)
	tutorService      services.TutorService      = services.NewTutorService(tutorRepository)

	accountController    controllers.AccountController    = controllers.NewAccountController(accountService)
	authCtrl             controllers.AuthController       = controllers.NewAuthController(authService, jwtService)
	subjectController    controllers.SubjectController    = controllers.NewSubjectController(subjectService)
	newClassController   controllers.NewClassController   = controllers.NewNewClassController(newClassService)
	classController      controllers.ClassController      = controllers.NewClassITController(classService)
	categoryController   controllers.CategoryController   = controllers.NewCategoryController(categoryService)
	postController       controllers.PostController       = controllers.NewPostController(postService)
	transController      controllers.TransController      = controllers.NewTransController(transService)
	salaryinfoController controllers.SalaryinfoController = controllers.NewSalaryinfoController(salaryinfoService)
	tutorController      controllers.TutorController      = controllers.NewTutorController(tutorService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "WELCOME TO GIASUANHEM ><"})
	})

	authRoutes := r.Group("v1/auth")
	{
		authRoutes.POST("/login", authCtrl.Login)
		authRoutes.POST("/login-admin", authCtrl.LoginAdmin)
		authRoutes.POST("/register", authCtrl.Register)
	}

	subjectRoutes := r.Group("v1/subject")
	{
		subjectRoutes.GET("/index", subjectController.FindAllSubject)
		subjectRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), subjectController.InsertSubject)
		subjectRoutes.GET("/id", subjectController.FindByID)
		subjectRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), subjectController.UpdateSubject)
		subjectRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), subjectController.DeleteSubject)
	}

	classRoutes := r.Group("v1/class")
	{
		classRoutes.GET("/index", classController.FindAllClass)
		classRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), classController.InsertClass)
		classRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), classController.DeleteClass)
		classRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), classController.UpdateClass)
		classRoutes.GET("/id", middleware.AuthorJWT(jwtService), classController.FindByID)
	}

	categoryRoutes := r.Group("v1/category")
	{
		categoryRoutes.GET("/index", categoryController.FindAllCategory)
		categoryRoutes.GET("/id", middleware.AuthorJWT(jwtService), categoryController.FindByID)
		categoryRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), categoryController.InsertCategory)
		categoryRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), categoryController.UpdateCategory)
		categoryRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), categoryController.DeleteCategory)
		categoryRoutes.GET("/filter", categoryController.FilterCategorry)
	}
	accountRoutes := r.Group("v1/account")
	{
		accountRoutes.GET("/index", accountController.FindAllAccount)
		accountRoutes.GET("/id", middleware.AuthorJWT(jwtService), accountController.FindByID)
		accountRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), accountController.DeleteAccount)
		accountRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), accountController.UpdateAccount)
		accountRoutes.GET("/filter",  accountController.FilterAccount)
		accountRoutes.POST("/password", middleware.AuthorJWTAdmin(jwtService), accountController.UpdatePassword)
	}

	newClassRoutes := r.Group("v1/new_class")
	{
		newClassRoutes.GET("/index", newClassController.FindAllNewClass)
		newClassRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), newClassController.InsertNewClass)
		newClassRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), newClassController.UpdateNewClass)
		newClassRoutes.GET("/id", middleware.AuthorJWT(jwtService), newClassController.FindByID)
		newClassRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), newClassController.DeleteNewClass)
		newClassRoutes.GET("/filter", newClassController.FilterNewClass)
		newClassRoutes.POST("/status", middleware.AuthorJWTAdmin(jwtService), newClassController.UpdateStatusNewClass)
	}

	postRoutes := r.Group("v1/post")
	{
		postRoutes.GET("/index", postController.FindAllPost)
		postRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), postController.InsertPost)
		postRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), postController.UpdatePost)
		postRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), postController.DeletePost)
		postRoutes.GET("/id", middleware.AuthorJWT(jwtService), postController.FindByID)
		postRoutes.GET("/filter", postController.FilterPost)
	}
	transRoutes := r.Group("v1/trans")
	{
		transRoutes.GET("/index",  transController.FindAllTrans)
		transRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), transController.InsertTrans)
		transRoutes.POST("/id", middleware.AuthorJWTAdmin(jwtService), transController.FindByIDAcc)
		transRoutes.GET("/filter", transController.FilterTrans)
		transRoutes.GET("/statistical", middleware.AuthorJWTAdmin(jwtService), transController.Statistics)
	}

	salRoutes := r.Group("v1/salaryinfo")
	{
		salRoutes.GET("/index", salaryinfoController.FindAllSalaryinfo)
		salRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), salaryinfoController.InsertSalaryinfo)
		salRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), salaryinfoController.DeleteSalaryinfo)
		salRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), salaryinfoController.UpdateSalaryinfo)
		salRoutes.GET("/id", middleware.AuthorJWT(jwtService), salaryinfoController.FindByID)
		salRoutes.GET("/filter", salaryinfoController.FindByType)
	}

	tutorRoutes := r.Group("v1/tutor")
	{
		tutorRoutes.GET("/index", tutorController.FindAllTutor)
		tutorRoutes.POST("/index", middleware.AuthorJWTAdmin(jwtService), tutorController.InsertTutor)
		tutorRoutes.GET("/id", middleware.AuthorJWT(jwtService), tutorController.FindByID)
		tutorRoutes.POST("/remove", middleware.AuthorJWTAdmin(jwtService), tutorController.DeleteTutor)
		tutorRoutes.POST("/edit", middleware.AuthorJWTAdmin(jwtService), tutorController.UpdateTutor)
		tutorRoutes.GET("/filter", tutorController.FilterTutor)
	}

	r.Run(":8100")
}
