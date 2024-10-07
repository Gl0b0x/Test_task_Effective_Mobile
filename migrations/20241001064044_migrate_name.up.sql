CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(50),
    song_name VARCHAR(50),
    release_date VARCHAR(10),
    text TEXT,
    link VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE (group_name, song_name)
);

INSERT INTO songs (group_name, song_name, release_date, text, link, created_at, updated_at)
VALUES
    ('Король и Шут', 'Кукла Колдуна', '15.05.1998',
     'Тёмный, мрачный коридор
Я на цыпочках, как вор
Пробираюсь, чуть дыша
Чтобы не спугнуть
Тех, кто спит уже давно
Тех, кому не всё равно
В чью я комнату тайком
Желаю заглянуть
Чтобы увидеть…

Как бессонница в час ночной
Меняет, нелюдимая, облик твой
Чьих, невольница, ты идей?
Зачем тебе охотиться на людей?

Крестик на моей груди
На него ты погляди
Что в тебе способен он
Резко изменить?
Много книжек я читал
Много фокусов видал
Свою тайну от меня
Не пытайся скрыть!
Я это видел!

Как бессонница в час ночной
Меняет, нелюдимая, облик твой
Чьих, невольница, ты идей?
Зачем тебе охотиться на людей?

Очень жаль, что ты тогда
Мне поверить не смогла
В то, что новый твой приятель
Не такой, как все!
Ты осталась с ним вдвоём
Не зная ничего о нём
Что для всех опасен он
Наплевать тебе!
И ты попала!

К настоящему колдуну
Он загубил таких, как ты, не одну!
Словно куклой, в час ночной
Теперь он может управлять тобой!

Всё происходит, будто в страшном сне
И находиться здесь опасно мне!',
     'https://www.youtube.com/watch?v=DHnJAjdrNdg', NOW(), NOW()
), ( 'Король и Шут', 'Прыгну со скалы', '14.03.1999',
    'C головы сорвал ветер мой колпак
Я хотел любви, но вышло всё не так
Знаю я, ничего в жизни не вернуть
И теперь у меня один лишь только путь

Разбежавшись, прыгну со скалы
Вот я был, и вот меня не стало
И когда об этом вдруг узнаешь ты
Тогда поймёшь, кого ты потеряла

Быть таким, как все, с детства не умел
Видимо, такой в жизни мой удел
А она, да что она? Вечно мне лгала
И меня никогда понять бы не смогла

Разбежавшись, прыгну со скалы
Вот я был, и вот меня не стало
И когда об этом вдруг узнаешь ты
Тогда поймёшь, кого ты потеряла

Гордо скину плащ, в даль направлю взор
Может, она ждёт? Вряд ли, это вздор
И, издав дикий крик, камнем брошусь вниз
Это моей жизни заключительный каприз

Разбежавшись, прыгну со скалы
Вот я был, и вот меня не стало
И тогда себя возненавидишь ты
Лишь осознав, кого ты потеряла
Кого ты потеряла
Кого ты потеряла',
    'https://genius.com/Korol-i-shut-jump-off-a-cliff-lyrics', NOW(), NOW()
), ('Король и Шут', 'Лесник', '20.10.1993',
'Замученный дорогой, я выбился из сил
И в доме лесника я ночлега попросил
С улыбкой добродушной старик меня впустил
И жестом дружелюбным на ужин пригласил
(Хэй!)

Будь как дома, путник
Я ни в чём не откажу
Я ни в чём не откажу
Я ни в чём не откажу! (Хэй!)
Множество историй
Коль желаешь, расскажу
Коль желаешь, расскажу
Коль желаешь, расскажу!

На улице темнело, сидел я за столом
Лесник сидел напротив, болтал о том, о сём
Что нет среди животных у старика врагов
Что нравится ему подкармливать волков

Будь как дома, путник
Я ни в чём не откажу
Я ни в чём не откажу
Я ни в чём не откажу! (Хэй!)
Множество историй
Коль желаешь, расскажу
Коль желаешь, расскажу
Коль желаешь, расскажу!

И волки среди ночи завыли под окном
Старик заулыбался и вдруг покинул дом
Но вскоре возвратился с ружьём наперевес
«Друзья хотят покушать, пойдём, приятель, в лес!»

Будь как дома, путник
Я ни в чём не откажу
Я ни в чём не откажу
Я ни в чём не откажу! (Хэй!)
Множество историй
Коль желаешь, расскажу
Коль желаешь, расскажу
Коль желаешь, расскажу!',
    'https://genius.com/Korol-i-shut-forester-lyrics', NOW(), NOW()
), ('muse', 'Supermassive Black Hole', '16.07.2006',
    'Ooh, baby, don''t you know I suffer?
Ooh, baby, can''t you hear me moan?
You caught me under false pretenses
How long before you let me go?

Ooh, you set my soul alight
Ooh, you set my soul alight

Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the supermassive (You set my soul alight)
Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the (You set my soul)
Into the supermassive

I thought I was a fool for no one
But, ooh, baby, I''m a fool for you
You''re the queen of the superficial
How long before you tell the truth?

Ooh, you set my soul alight
Ooh, you set my soul alight

Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the supermassive (You set my soul alight)
Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the (You set my soul)
Into the supermassive

Supermassive black hole
Supermassive black hole
Supermassive black hole
Supermassive black hole

Glaciers melting in the dead of night
And the superstar''s sucked into the supermassive
Glaciers melting in the dead of night
And the superstar''s sucked into the supermassive
Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the supermassive (You set my soul alight)
Glaciers melting in the dead of night (Ooh)
And the superstar''s sucked into the (You set my soul)
Into the supermassive

Supermassive black hole
Supermassive black hole
Supermassive black hole
Supermassive black hole',
    'https://www.youtube.com/watch?v=Xsp3_a-PMTw',NOW(), NOW()
), (
    'muse', 'Starlight', '16.07.2006',
    'Far away
This ship is taking me far away
Far away from the memories
Of the people who care if I live or die

Starlight
I will be chasing a starlight
Until the end of my life
I don''t know if it''s worth it anymore

Hold you in my arms
I just wanted to hold you in my arms

My life
You electrify my life
Let''s conspire to ignite
All the souls that would die just to feel alive

I''ll never let you go
If you promise not to fade away
Never fade away

And our hopes and expectations
Black holes and revelations
And our hopes and expectations
Black holes and revelations

Hold you in my arms
I just wanted to hold you in my arms

Far away
This ship has taken me far away
Far away from the memories
Of the people who care if I live or die

I''ll never let you go
If you promise not to fade away
Never fade away

And our hopes and expectations
Black holes and revelations, yeah
Our hopes and expectations
Black holes and revelations

Hold you in my arms
I just wanted to hold you in my arms
I just wanted to hold',
    'https://www.youtube.com/watch?v=Xsp3_a-PMTw',NOW(), NOW()
);