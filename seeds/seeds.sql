-- Seed work experiences
INSERT INTO work_experiences (company_name, position, description, start_date, end_date) VALUES
    ('Tech Solutions Inc.', 'Senior Software Engineer', 'Led development of cloud-based applications using microservices architecture. Mentored junior developers and implemented CI/CD pipelines.', '2020-01', '2023-12'),
    ('Digital Innovations LLC', 'Full Stack Developer', 'Developed and maintained multiple web applications using React and Node.js. Improved application performance by 40%.', '2018-03', '2019-12'),
    ('StartUp Hub', 'Junior Developer', 'Collaborated in an agile team to build responsive web applications. Implemented user authentication and authorization systems.', '2017-06', '2018-02');

-- Seed project experiences
INSERT INTO project_experiences (title, description, technologies, github_link, live_link) VALUES
    ('E-commerce Platform', 'A full-stack e-commerce solution with shopping cart and payment integration.', 'React, Node.js, PostgreSQL, Stripe', 'https://github.com/username/ecommerce', 'https://demo-ecommerce.com'),
    ('Task Management App', 'Kanban-style project management tool with real-time updates.', 'Vue.js, Express, MongoDB, Socket.io', 'https://github.com/username/task-manager', 'https://task-app-demo.com'),
    ('Weather Dashboard', 'Weather forecast application with location-based services.', 'JavaScript, OpenWeather API, HTML5, CSS3', 'https://github.com/username/weather-app', 'https://weather-dashboard-demo.com');