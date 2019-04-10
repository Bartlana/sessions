/******************Предметы******************/

insert into subjects (subject_name) values ('Zalupa');
insert into subjects (subject_name) values ('Polnaya Zalupa');
insert into subjects (subject_name) values ('Об''єктно-орiєнтоване програмування');
insert into subjects (subject_name) values ('Структурне та логічне програмування');
insert into subjects (subject_name) values ('Процедурне програмування');
insert into subjects (subject_name) values ('Інформаційно-аналітична діяльність у галузі інформаційної безпеки');
insert into subjects (subject_name) values ('Грід-системи та технології хмарних обчислень');
insert into subjects (subject_name) values ('Дискретна математика');
insert into subjects (subject_name) values ('Технічна термодинаміка');
insert into subjects (subject_name) values ('Сталий розвитокк дестинацій: науково-технічний аспект');
insert into subjects (subject_name) values ('Застосування сонячної енергiї');
insert into subjects (subject_name) values ('Фізико-хімічні основи процесів перетворення енергії ');
insert into subjects (subject_name) values ('Управління якістю продукції у виробництві');
insert into subjects (subject_name) values ('Метрологія та стандартизація ');
insert into subjects (subject_name) values ('Основи метрології, стандартизації та контролю якості');
insert into subjects (subject_name) values ('Конструкція, міцність та надійність газотурбінних установок');
insert into subjects (subject_name) values ('Суднова холодильна техніка');
insert into subjects (subject_name) values ('Процеси та елементна база систем нетрадиційної енергетики');
insert into subjects (subject_name) values ('Електротехніка та електромеханіка');
insert into subjects (subject_name) values ('Фізика');

/******************Преподаватели******************/

insert into professors (professor_name, login, password)
values ('Пупкин Василий Иванович', 'Пупкин Василий Иванович', 'Пупкин Василий Иванович');
insert into professors (professor_name, login, password)
values ('Плотніков Валерій Михайлович', 'Плотніков Валерій Михайлович', 'Плотніков Валерій Михайлович');
insert into professors (professor_name, login, password)
values ('Снігур Тетяна Сергіївна', 'Снігур Тетяна Сергіївна', 'Снігур Тетяна Сергіївна');
insert into professors (professor_name, login, password)
values ('Сіромля Сергій Григорович', 'Сіромля Сергій Григорович', 'Сіромля Сергій Григорович');
insert into professors (professor_name, login, password)
values ('Монтик Петр Николаевич', 'Монтик Петр Николаевич', 'Монтик Петр Николаевич');
insert into professors (professor_name, login, password)
values ('Галиулин Анатолий Агзамович', 'Галиулин Анатолий Агзамович', 'Галиулин Анатолий Агзамович');

/******************Группы******************/

insert into groups (group_name) values ('y1');
insert into groups (group_name) values ('345');
insert into groups (group_name) values ('215');
insert into groups (group_name) values ('kk5');
insert into groups (group_name) values ('321');



/******************Занятия******************/

insert into class (class_date, subject, professor)
values ('03-09-2019', 1, 1);
insert into class (class_date, subject, professor)
values ('04-09-2019', 2, 1);
insert into class (class_date, subject, professor)
values ('05-09-2019', 1, 1);
insert into class (class_date, subject, professor)
values ('06-09-2019', 2, 1);
insert into class (class_date, subject, professor)
values ('07-09-2019', 1, 1);
insert into class (class_date, subject, professor)
values ('08-09-2019', 2, 1);
insert into class (class_date, subject, professor)
values ('09-09-2019', 1, 1);
insert into class (class_date, subject, professor)
values ('10-09-2019', 2, 1);
insert into class (class_date, subject, professor)
values ('11-09-2019', 1, 1);
insert into class (class_date, subject, professor)
values ('12-09-2019', 2, 1);

/******************Студенты******************/

insert into students (student_name, group_id) values ('Білий Володимир Сергійович', 5);
insert into students (student_name, group_id) values ('Вінниченко Наталя Олександрівна', 1);
insert into students (student_name, group_id) values ('Волкова Анастасія Юріївна', 2);
insert into students (student_name, group_id) values ('Голубніченко Максим Андрійович', 3);
insert into students (student_name, group_id) values ('Гриб Максим Григорович', 4);
insert into students (student_name, group_id) values ('Дерменжи Тетяна Степанівна', 5);
insert into students (student_name, group_id) values ('Диков Олексій Сергійович', 5);
insert into students (student_name, group_id) values ('Кожин Сергій Дмитрович', 1);
insert into students (student_name, group_id) values ('Круглей Ольга Володимирівна', 2);
insert into students (student_name, group_id) values ('Нікорчук Олександр Валерійович', 3);
insert into students (student_name, group_id) values ('Осіпов Георгій Едуардович', 4);
insert into students (student_name, group_id) values ('Прус Вікторія Вячеславівна', 5);
insert into students (student_name, group_id) values ('Рамах Юссеф Глі Терсаі', 5);
insert into students (student_name, group_id) values ('Соловйов Едуард Григорійович', 1);
insert into students (student_name, group_id) values ('Харахаш Олександр Вячеславович', 2);
insert into students (student_name, group_id) values ('Чараєв Костянтин Юрійович', 3);
insert into students (student_name, group_id) values ('Барткова Світлана Олександрівна', 4);
insert into students (student_name, group_id) values ('Безносюк Людмила Олександрівна', 5);
insert into students (student_name, group_id) values ('Борщьов Дмитро Вікторович', 5);
insert into students (student_name, group_id) values ('Виходченко Володимир Анатолійович', 1);
insert into students (student_name, group_id) values ('Дударець Катерина Вадимівна', 4);
insert into students (student_name, group_id) values ('Іванова Христина Владиславівна', 2);
insert into students (student_name, group_id) values ('Крушельницький Владислав Олегович', 3);
insert into students (student_name, group_id) values ('Лисенко Єлизавета Сергіївна', 4);
insert into students (student_name, group_id) values ('Лукіянчук Юрій Андрійович', 5);
insert into students (student_name, group_id) values ('Лятанська Валерія Олегівна', 5);
insert into students (student_name, group_id) values ('Орехова Анастасія Ігорівна', 1);
insert into students (student_name, group_id) values ('Орешет Вікторія Сергіївна', 2);
insert into students (student_name, group_id) values ('Помогаєв Дмитро Олексійович', 3);
insert into students (student_name, group_id) values ('Приймак Олександр Валерійович', 4);
insert into students (student_name, group_id) values ('Таланов Ілля Вячеславович', 5);
insert into students (student_name, group_id) values ('Баранов Мішель Михайлович', 5);
insert into students (student_name, group_id) values ('Бєлоус Владислав Олегович', 1);
insert into students (student_name, group_id) values ('Бужор Владислав Анатолійович', 2);
insert into students (student_name, group_id) values ('Гавриленко Денис Олегович', 3);
insert into students (student_name, group_id) values ('Добряк Леонід Анатолійович', 4);
insert into students (student_name, group_id) values ('Кічук Ігор Леонідович', 5);
insert into students (student_name, group_id) values ('Літвіненко Георгій Ігорович', 5);
insert into students (student_name, group_id) values ('Мемрук Дмитро Сергійович', 1);
insert into students (student_name, group_id) values ('Пінтя Вадим Сергійович', 2);
insert into students (student_name, group_id) values ('Погребняк Олег Вячеславович', 3);
insert into students (student_name, group_id) values ('Соколи Іван Петрович', 4);
insert into students (student_name, group_id) values ('Сотніченко Максим Сергійович', 5);
insert into students (student_name, group_id) values ('Ханчевський Владислав Андрійович', 5);
insert into students (student_name, group_id) values ('Вдовиченко Максим Ігорович', 1);
insert into students (student_name, group_id) values ('Гуменюк Микола Віталійович', 2);
insert into students (student_name, group_id) values ('Димов Олександр Вячеславович', 3);
insert into students (student_name, group_id) values ('Єнов Богдан Олександрович', 3);
insert into students (student_name, group_id) values ('Єпур Лілія Ігорівна', 4);
insert into students (student_name, group_id) values ('Іоргачов Євген Юрійович', 5);
insert into students (student_name, group_id) values ('Каднова Аліна Валеріївна', 5);
insert into students (student_name, group_id) values ('Луценко Артем Геннадійович', 5);
insert into students (student_name, group_id) values ('Степул Артем Мартиросович', 1);
insert into students (student_name, group_id) values ('Ткачук Станіслав Віталійович', 1);
insert into students (student_name, group_id) values ('Хлюстов Юрій Віталійович', 2);
insert into students (student_name, group_id) values ('Цибуляк Дмитрій Володимирович', 2);
insert into students (student_name, group_id) values ('Шевченко Валентин Павлович', 5);


/******************Вопросы******************/



/******************Ответы******************/



/******************Присутствие******************/