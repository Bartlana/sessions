create table groups (
  group_id     bigserial primary key,
  group_name   text not null,
  subgroup     text
);

create table professors (
  professor_id    bigserial primary key,
  professor_name  text not null,
  login           text not null,
  password        text not null
);

create table subjects (
  subject_id    bigserial primary key,
  subject_name  text not null
);

create table students (
  student_id      bigserial primary key,
  student_name    text      not null,
  group_id        bigint    references groups(group_id),
  login           text      not null,
  password        text      not null
);

create table questions (
  question_id      bigserial primary key,
  question_text    text not null
);

create table answers (
  answer_id   bigserial primary key,
  question    bigint references questions(question_id),
  answer_text text not null,
  student     bigint references students(student_id)
);

create table statistic (
  statistic_id bigserial primary key,
  question     bigint references questions(question_id),
  student      bigint references students(student_id)
);

create table class (
  class_id    bigserial primary key,
  class_date  date not null,
  subject     bigint references subjects(subject_id),
  professor   bigint references professors(professor_id),
  theme       text
);

create table group_in_class (
  class_id bigint references class(class_id),
  group_id bigint references groups(group_id)
);

create table presence (
  presence_id bigserial primary key,
  class       bigint    references class(class_id),
  student     bigint    references students(student_id)
);