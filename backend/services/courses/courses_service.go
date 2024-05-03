package courses

import coursesDomain "backend/domain/courses"

func CreateCourse(request coursesDomain.CourseCreateRequest) coursesDomain.CreateCourseResponse {

	//create in db

	return coursesDomain.CreateCourseResponse{1, request.Name, request.Description, request.Price, request.ClassAmount, request.StartDate, request.EndDate}
}
