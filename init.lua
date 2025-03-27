box.schema.user.create('vote_bot', {password = 'password'})
box.schema.user.grant('vote_bot', 'read,write,execute', 'universe')

-- Создание пространств
polls = box.schema.space.create('polls', {
    if_not_exists = true,
    format = {
        {name = 'id', type = 'string'},
        {name = 'creator', type = 'string'},
        {name = 'question', type = 'string'},
        {name = 'options', type = 'map'},
        {name = 'status', type = 'string'}, -- 'active'/'closed'
        {name = 'created_at', type = 'unsigned'}
    },
    indexes = {
        {name = 'primary', type = 'hash', parts = {'id'}},
        {name = 'creator_idx', type = 'tree', parts = {'creator'}}
    }
})

votes = box.schema.space.create('votes', {
    if_not_exists = true,
    format = {
        {name = 'poll_id', type = 'string'},
        {name = 'user_id', type = 'string'},
        {name = 'option', type = 'string'},
        {name = 'voted_at', type = 'unsigned'}
    },
    indexes = {
        {name = 'primary', type = 'tree', parts = {'poll_id', 'user_id'}},
        {name = 'poll_idx', type = 'tree', parts = {'poll_id'}}
    }
})