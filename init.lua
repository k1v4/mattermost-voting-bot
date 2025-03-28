box.cfg{
    listen = 3301
}

if not box.schema.user.exists('vote_bot') then
    box.schema.user.create('vote_bot', { password = 'password' })
end

box.schema.user.grant('vote_bot', 'read,write,execute', 'universe')

polls = box.schema.space.create('polls', { if_not_exists = true })

polls:format({
    {name = 'id', type = 'string'},
    {name = 'creator', type = 'string'},
    {name = 'question', type = 'string'},
    {name = 'options', type = 'map'},
    {name = 'status', type = 'string'}, -- 'active'/'closed'
    {name = 'created_at', type = 'unsigned'}
})

polls:create_index('primary', {
    type = 'hash',
    parts = {'id'}
})

polls:create_index('creator_idx', {
    type = 'tree',
    parts = {'creator'}
})


votes = box.schema.space.create('votes', { if_not_exists = true })

votes:format({
    {name = 'poll_id', type = 'string'},
    {name = 'user_id', type = 'string'},
    {name = 'option', type = 'string'},
    {name = 'voted_at', type = 'unsigned'}
})

votes:create_index('primary', {
    type = 'tree',
    parts = {'poll_id', 'user_id'}
})

votes:create_index('poll_idx', {
    type = 'tree',
    parts = {'poll_id'}
})

