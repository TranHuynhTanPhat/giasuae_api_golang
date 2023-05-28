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
		authRoutes.POST("/register", authCtrl.Register)
	}

	subjectRoutes := r.Group("v1/subject")
	{
		subjectRoutes.GET("/index", subjectController.FindAllSubject)
		subjectRoutes.POST("/index", middleware.AuthorJWT(jwtService), subjectController.InsertSubject)
		subjectRoutes.GET("/id", subjectController.FindByID)
		subjectRoutes.POST("/edit", middleware.AuthorJWT(jwtService), subjectController.UpdateSubject)
		subjectRoutes.POST("/remove", middleware.AuthorJWT(jwtService), subjectController.DeleteSubject)
	}

	classRoutes := r.Group("v1/class")
	{
		classRoutes.GET("/index", middleware.AuthorJWT(jwtService), classController.FindAllClass)
		classRoutes.POST("/index", middleware.AuthorJWT(jwtService), classController.InsertClass)
		classRoutes.POST("/remove", middleware.AuthorJWT(jwtService), classController.DeleteClass)
		classRoutes.POST("/edit", middleware.AuthorJWT(jwtService), classController.UpdateClass)
		classRoutes.GET("/id", middleware.AuthorJWT(jwtService), classController.FindByID)
	}

	categoryRoutes := r.Group("v1/category")
	{
		categoryRoutes.GET("/index", middleware.AuthorJWT(jwtService), categoryController.FindAllCategory)
		categoryRoutes.GET("/id", middleware.AuthorJWT(jwtService), categoryController.FindByID)
		categoryRoutes.POST("/index", middleware.AuthorJWT(jwtService), categoryController.InsertCategory)
		categoryRoutes.POST("/edit", middleware.AuthorJWT(jwtService), categoryController.UpdateCategory)
		categoryRoutes.POST("/remove", middleware.AuthorJWT(jwtService), categoryController.DeleteCategory)
		categoryRoutes.GET("/filter",middleware.AuthorJWT(jwtService), categoryController.FilterCategorry)
	}
	accountRoutes := r.Group("v1/account")
	{
		accountRoutes.GET("/index", middleware.AuthorJWT(jwtService), accountController.FindAllAccount)
		accountRoutes.GET("/id", middleware.AuthorJWT(jwtService), accountController.FindByID)
		accountRoutes.POST("/remove", middleware.AuthorJWT(jwtService), accountController.DeleteAccount)
		accountRoutes.POST("/edit", middleware.AuthorJWT(jwtService), accountController.UpdateAccount)
		accountRoutes.GET("/filter", middleware.AuthorJWT(jwtService), accountController.FilterAccount)
		accountRoutes.POST("/password", middleware.AuthorJWT(jwtService), accountController.UpdatePassword)
	}

	newClassRoutes := r.Group("v1/new_class")
	{
		newClassRoutes.GET("/index", middleware.AuthorJWT(jwtService), newClassController.FindAllNewClass)
		newClassRoutes.POST("/index", middleware.AuthorJWT(jwtService), newClassController.InsertNewClass)
		newClassRoutes.POST("/edit", middleware.AuthorJWT(jwtService), newClassController.UpdateNewClass)
		newClassRoutes.GET("/id", middleware.AuthorJWT(jwtService), newClassController.FindByID)
		newClassRoutes.POST("/remove", middleware.AuthorJWT(jwtService), newClassController.DeleteNewClass)
		newClassRoutes.GET("/filter", middleware.AuthorJWT(jwtService), newClassController.FilterNewClass)
		newClassRoutes.POST("/status", middleware.AuthorJWT(jwtService), newClassController.UpdateStatusNewClass)
	}

	postRoutes := r.Group("v1/post")
	{
		postRoutes.GET("/index", middleware.AuthorJWT(jwtService), postController.FindAllPost)
		postRoutes.POST("/index", middleware.AuthorJWT(jwtService), postController.InsertPost)
		postRoutes.POST("/edit", middleware.AuthorJWT(jwtService), postController.UpdatePost)
		postRoutes.POST("/remove", middleware.AuthorJWT(jwtService), postController.DeletePost)
		postRoutes.GET("/id", middleware.AuthorJWT(jwtService), postController.FindByID)
		postRoutes.GET("/filter", middleware.AuthorJWT(jwtService), postController.FilterPost)
	}
	transRoutes := r.Group("v1/trans")
	{
		transRoutes.GET("/index", middleware.AuthorJWT(jwtService), transController.FindAllTrans)
		transRoutes.POST("/index", middleware.AuthorJWT(jwtService), transController.InsertTrans)
		transRoutes.POST("/id", middleware.AuthorJWT(jwtService), transController.FindByIDAcc)
		transRoutes.GET("/filter", middleware.AuthorJWT(jwtService), transController.FilterTrans)
		transRoutes.GET("/statistical", middleware.AuthorJWT(jwtService), transController.Statistics)
	}

	salRoutes := r.Group("v1/salaryinfo")
	{
		salRoutes.GET("/index", middleware.AuthorJWT(jwtService), salaryinfoController.FindAllSalaryinfo)
		salRoutes.POST("/index", middleware.AuthorJWT(jwtService), salaryinfoController.InsertSalaryinfo)
		salRoutes.POST("/remove", middleware.AuthorJWT(jwtService), salaryinfoController.DeleteSalaryinfo)
		salRoutes.POST("/edit", middleware.AuthorJWT(jwtService), salaryinfoController.UpdateSalaryinfo)
		salRoutes.GET("/id", middleware.AuthorJWT(jwtService), salaryinfoController.FindByID)
		salRoutes.GET("/filter", middleware.AuthorJWT(jwtService), salaryinfoController.FindByType)
	}

	tutorRoutes := r.Group("v1/tutor")
	{
		tutorRoutes.GET("/index", middleware.AuthorJWT(jwtService), tutorController.FindAllTutor)
		tutorRoutes.POST("/index", middleware.AuthorJWT(jwtService), tutorController.InsertTutor)
		tutorRoutes.GET("/id", middleware.AuthorJWT(jwtService), tutorController.FindByID)
		tutorRoutes.POST("/remove", middleware.AuthorJWT(jwtService), tutorController.DeleteTutor)
		tutorRoutes.POST("/edit", middleware.AuthorJWT(jwtService), tutorController.UpdateTutor)
		tutorRoutes.GET("/filter", middleware.AuthorJWT(jwtService), tutorController.FilterTutor)
	}

	r.Run(":8100")
}
