package courses

import coursesDomain "backend/domain/courses"

func CreateCourse(request coursesDomain.CourseCreateRequest) coursesDomain.CreateCourseResponse {

	//create in db

	return coursesDomain.CreateCourseResponse{ID: 1, Name: request.Name, Description: request.Description, Price: request.Price, ClassAmount: request.ClassAmount, StartDate: request.StartDate, EndDate: request.EndDate}
}
