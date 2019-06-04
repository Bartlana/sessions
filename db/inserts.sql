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

insert into classes (class_date, subject, professor)
values ('03-09-2019', 4, 1);
insert into classes (class_date, subject, professor)
values ('04-09-2019', 5, 1);
insert into classes (class_date, subject, professor)
values ('05-09-2019', 4, 1);
insert into classes (class_date, subject, professor)
values ('06-09-2019', 5, 1);
insert into classes (class_date, subject, professor)
values ('07-09-2019', 4, 1);
insert into classes (class_date, subject, professor)
values ('08-09-2019', 5, 1);
insert into classes (class_date, subject, professor)
values ('09-09-2019', 4, 1);
insert into classes (class_date, subject, professor)
values ('10-09-2019', 5, 1);
insert into classes (class_date, subject, professor)
values ('11-09-2019', 4, 1);
insert into classes (class_date, subject, professor)
values ('12-09-2019', 5, 1);
insert into classes (class_date, subject, professor)
values ('12-09-2019', 14, 3);
insert into classes (class_date, subject, professor)
values ('12-09-2019', 14, 3);
insert into classes (class_date, subject, professor)
values ('12-09-2019', 14, 3);
insert into classes (class_date, subject, professor)
values ('12-09-2019', 14, 3);

insert into group_in_class (class_id, group_id)
values (3, 1);
insert into group_in_class (class_id, group_id)
values (3, 2);
insert into group_in_class (class_id, group_id)
values (3, 3);
insert into group_in_class (class_id, group_id)
values (3, 4);
insert into group_in_class (class_id, group_id)
values (4, 1);
insert into group_in_class (class_id, group_id)
values (4, 9);
insert into group_in_class (class_id, group_id)
values (4, 8);
insert into group_in_class (class_id, group_id)
values (4, 7);
insert into group_in_class (class_id, group_id)
values (5, 5);
insert into group_in_class (class_id, group_id)
values (5, 1);
insert into group_in_class (class_id, group_id)
values (6, 5);
insert into group_in_class (class_id, group_id)
values (7, 1);
insert into group_in_class (class_id, group_id)
values (7, 4);
insert into group_in_class (class_id, group_id)
values (7, 7);
insert into group_in_class (class_id, group_id)
values (8, 1);
insert into group_in_class (class_id, group_id)
values (9, 4);
insert into group_in_class (class_id, group_id)
values (10, 4);
insert into group_in_class (class_id, group_id)
values (10, 5);
insert into group_in_class (class_id, group_id)
values (13, 1);
insert into group_in_class (class_id, group_id)
values (14, 1);
insert into group_in_class (class_id, group_id)
values (15, 1);
insert into group_in_class (class_id, group_id)
values (16, 1);

/******************Студенты******************/

insert into students (student_name, group_id, login, password) values ('Білий Володимир Сергійович', 5, '123', '123');
insert into students (student_name, group_id, login, password) values ('Вінниченко Наталя Олександрівна', 1, '1243', '123');
insert into students (student_name, group_id, login, password) values ('Волкова Анастасія Юріївна', 2, '1234', '123');
insert into students (student_name, group_id, login, password) values ('Голубніченко Максим Андрійович', 3, '12345', '123');
insert into students (student_name, group_id, login, password) values ('Гриб Максим Григорович', 4, '123456', '123');
insert into students (student_name, group_id, login, password) values ('Дерменжи Тетяна Степанівна', 5, '1234567', '123');
insert into students (student_name, group_id, login, password) values ('Диков Олексій Сергійович', 5, '12345678', '123');
insert into students (student_name, group_id, login, password) values ('Кожин Сергій Дмитрович', 1, '123456789', '123');
insert into students (student_name, group_id, login, password) values ('Круглей Ольга Володимирівна', 2, '12m5nyt30', '123');
insert into students (student_name, group_id, login, password) values ('Нікорчук Олександр Валерійович', 3, '1239', '123');
insert into students (student_name, group_id, login, password) values ('Осіпов Георгій Едуардович', 4, '123mten8', '123');
insert into students (student_name, group_id, login, password) values ('Прус Вікторія Вячеславівна', 5, '1237', '123');
insert into students (student_name, group_id, login, password) values ('Рамах Юссеф Глі Терсаі', 5, '12 mznb36', '123');
insert into students (student_name, group_id, login, password) values ('Соловйов Едуард Григорійович', 1, '12uymn35', '123');
insert into students (student_name, group_id, login, password) values ('Харахаш Олександр Вячеславович', 2, '1723', '123');
insert into students (student_name, group_id, login, password) values ('Чараєв Костянтин Юрійович', 3, '11lyfk23', '123');
insert into students (student_name, group_id, login, password) values ('Барткова Світлана Олександрівна', 4, '21idu23', '123');
insert into students (student_name, group_id, login, password) values ('Безносюк Людмила Олександрівна', 5, '3mny5tj123', '123');
insert into students (student_name, group_id, login, password) values ('Борщьов Дмитро Вікторович', 5, '41xctgkj23', '123');
insert into students (student_name, group_id, login, password) values ('Виходченко Володимир Анатолійович', 1, '5123', '123');
insert into students (student_name, group_id, login, password) values ('Дударець Катерина Вадимівна', 4, '6123', '123');
insert into students (student_name, group_id, login, password) values ('Іванова Христина Владиславівна', 2, '7123', '123');
insert into students (student_name, group_id, login, password) values ('Крушельницький Владислав Олегович', 3, '8123', '123');
insert into students (student_name, group_id, login, password) values ('Лисенко Єлизавета Сергіївна', 4, '9123', '123');
insert into students (student_name, group_id, login, password) values ('Лукіянчук Юрій Андрійович', 5, '01smn23', '123');
insert into students (student_name, group_id, login, password) values ('Лятанська Валерія Олегівна', 5, '1213', '123');
insert into students (student_name, group_id, login, password) values ('Орехова Анастасія Ігорівна', 1, '12m5nyset23', '123');
insert into students (student_name, group_id, login, password) values ('Орешет Вікторія Сергіївна', 2, '12m5nye33', '123');
insert into students (student_name, group_id, login, password) values ('Помогаєв Дмитро Олексійович', 3, '12jn43', '123');
insert into students (student_name, group_id, login, password) values ('Приймак Олександр Валерійович', 4, '12573', '123');
insert into students (student_name, group_id, login, password) values ('Таланов Ілля Вячеславович', 5, '1263', '123');
insert into students (student_name, group_id, login, password) values ('Баранов Мішель Михайлович', 5, '1273', '123');
insert into students (student_name, group_id, login, password) values ('Бєлоус Владислав Олегович', 1, '1283', '123');
insert into students (student_name, group_id, login, password) values ('Бужор Владислав Анатолійович', 2, '1jlkj123', '123');
insert into students (student_name, group_id, login, password) values ('Гавриленко Денис Олегович', 3, '1b[ni223', '123');
insert into students (student_name, group_id, login, password) values ('Добряк Леонід Анатолійович', 4, '1olukg323', '123');
insert into students (student_name, group_id, login, password) values ('Кічук Ігор Леонідович', 5, '13sm7n6423', '123');
insert into students (student_name, group_id, login, password) values ('Літвіненко Георгій Ігорович', 5, '152s7mnr4y3', '123');
insert into students (student_name, group_id, login, password) values ('Мемрук Дмитро Сергійович', 1, '156wsmj23', '123');
insert into students (student_name, group_id, login, password) values ('Пінтя Вадим Сергійович', 2, '1smk57n4jh7823', '123');
insert into students (student_name, group_id, login, password) values ('Погребняк Олег Вячеславович', 3, '1ryh23', '123');
insert into students (student_name, group_id, login, password) values ('Соколи Іван Петрович', 4, '12gj3', '123');
insert into students (student_name, group_id, login, password) values ('Сотніченко Максим Сергійович', 5, '1jtgu23', '123');
insert into students (student_name, group_id, login, password) values ('Ханчевський Владислав Андрійович', 5, '12ki3', '123');
insert into students (student_name, group_id, login, password) values ('Вдовиченко Максим Ігорович', 1, '1gitjn23', '123');
insert into students (student_name, group_id, login, password) values ('Гуменюк Микола Віталійович', 2, '12hlj3', '123');
insert into students (student_name, group_id, login, password) values ('Димов Олександр Вячеславович', 3, '1vkm23', '123');
insert into students (student_name, group_id, login, password) values ('Єнов Богдан Олександрович', 3, '123bhkmuj', '123');
insert into students (student_name, group_id, login, password) values ('Єпур Лілія Ігорівна', 4, '123ikjh', '123');
insert into students (student_name, group_id, login, password) values ('Іоргачов Євген Юрійович', 5, '12b vhjm3', '123');
insert into students (student_name, group_id, login, password) values ('Каднова Аліна Валеріївна', 5, '12vjg3', '123');
insert into students (student_name, group_id, login, password) values ('Луценко Артем Геннадійович', 5, '1zyetgdf23', '123');
insert into students (student_name, group_id, login, password) values ('Степул Артем Мартиросович', 1, '1ckyhv23', '123');
insert into students (student_name, group_id, login, password) values ('Ткачук Станіслав Віталійович', 1, '12ckyjh3', '123');
insert into students (student_name, group_id, login, password) values ('Хлюстов Юрій Віталійович', 2, '12ckyh3', '123');
insert into students (student_name, group_id, login, password) values ('Цибуляк Дмитрій Володимирович', 2, '1to8i23', '123');
insert into students (student_name, group_id, login, password) values ('Шевченко Валентин Павлович', 5, '12y4rxg3', '123');


/******************Вопросы******************/



/******************Ответы******************/



/******************Присутствие******************/
insert into presences (class, student) values (3, 100);
insert into presences (class, student) values (3, 101);
insert into presences (class, student) values (3, 102);
insert into presences (class, student) values (3, 103);
insert into presences (class, student) values (3, 104);
insert into presences (class, student) values (3, 105);
insert into presences (class, student) values (3, 106);
insert into presences (class, student) values (3, 107);


insert into presences (class, student) values (4, 100);
insert into presences (class, student) values (4, 101);
insert into presences (class, student) values (4, 102);
insert into presences (class, student) values (4, 103);
insert into presences (class, student) values (4, 104);
insert into presences (class, student) values (4, 105);
insert into presences (class, student) values (4, 106);
insert into presences (class, student) values (4, 107);

insert into presences (class, student) values (4, 110);
insert into presences (class, student) values (4, 111);
insert into presences (class, student) values (4, 112);
insert into presences (class, student) values (4, 113);
insert into presences (class, student) values (4, 114);
insert into presences (class, student) values (4, 115);
insert into presences (class, student) values (4, 116);
insert into presences (class, student) values (4, 117);

insert into presences (class, student) values (5, 140);
insert into presences (class, student) values (5, 141);
insert into presences (class, student) values (5, 142);
insert into presences (class, student) values (5, 143);
insert into presences (class, student) values (5, 144);
insert into presences (class, student) values (5, 145);
insert into presences (class, student) values (5, 146);
insert into presences (class, student) values (5, 147);