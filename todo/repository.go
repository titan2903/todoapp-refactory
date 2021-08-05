package todo

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Todo, error)
	FindByUserID(userID int) ([]Todo, error)
	FindByUserIDAndTodoId(ID int, userID int) (Todo, error)
	FindByID(ID int) (Todo, error)
	SaveTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodo(todo Todo) (Todo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository { //! membuat object baru dari repository dan nilai db dari repository di isi sesuai parameter di NewRepository
	return &repository{db}
}

func(r *repository) FindAll() ([]Todo, error) {
	var todos []Todo

	err := r.db.Find(&todos).Error //! type harus pointer
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func(r *repository) FindByUserID(userID int) ([]Todo, error) {
	var todos []Todo

	err := r.db.Where("user_id = ?", userID).Find(&todos).Error //! preload merupakan ngeload relasinya dan mengambil data relasinya
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func(r *repository) FindByUserIDAndTodoId(ID int, userID int) (Todo, error) {
	var todo Todo

	err := r.db.Where("user_id = ? AND id = ?", userID, ID).Find(&todo).Error //! preload merupakan ngeload relasinya dan mengambil data relasinya
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func(r *repository) FindByID(ID int) (Todo, error) {
	var todo Todo

	err := r.db.Preload("User").Where("id = ?", ID).Find(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func(r *repository) SaveTodo(todo Todo) (Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func(r *repository) UpdateTodo(todo Todo) (Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func(r *repository) DeleteTodo(todo Todo) (Todo, error) {
	err := r.db.Delete(&todo).Error
	if err != nil {
		return todo, err
	}

	return todo, nil
}