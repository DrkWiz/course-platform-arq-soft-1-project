import React from 'react';
import Cookies from 'js-cookie';
import CourseBlock from './CourseBlock'; // Asegúrate de importar el componente CourseBlock

class MyCourses extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      courses: [],
    };
  }

  componentDidMount() {
    const token = Cookies.get('token'); // Obtiene el token de la cookie

    fetch('http://localhost:8080/users/courses/:id', {
      headers: {
        'Authorization': `Bearer ${token}` // Envia el token en el header de la solicitud
      }
    })
      .then(response => response.json())
      .then(data => this.setState({ courses: data }));
  }

  render() {
    return (
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {this.state.courses.map((course) => (
          <div key={course.id} style={{ flex: '0 0 50%' }}>
            <CourseBlock course={course} />
          </div>
        ))}
      </div>
    );
  }
}

export default MyCourses;