-- GetTeacherDisciplinesAndClasses method
SELECT JSON_BUILD_OBJECT(
    'id', sg.discipline_id,
    'title', (SELECT name FROM discipline dis WHERE dis.id = sg.discipline_id),
    'classes', (
        SELECT JSON_AGG(
            JSON_BUILD_OBJECT(
                'id', c.id,
                'title', concat(c.grade, ' ', c.title),
                'main_teacher_id', c.main_teacher_id,
                'students', (
                    SELECT JSON_AGG(
                        JSON_BUILD_OBJECT(
                            'id', s.id,
                            'first_name', s.first_name,
                            'middle_name', s.middle_name,
                            'last_name', s.last_name
                        )
                    ) FROM users s WHERE s.classes_id = c.id AND s.role_id = 1
                )
            )
        ) FROM classes c
    )
) AS result
FROM study_group sg
WHERE teacher_id = $1;