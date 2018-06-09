package gache

const setScript = `
redis.call('setex', ARGV[1], ARGV[3], ARGV[2])

for i, tag in ipairs(ARGV) do
	if i > 3 then
		redis.call('sadd', tag .. '.tags', ARGV[1])
	end
end
`

const bustScript = `
for _, tag in ipairs(ARGV) do
	for _, key in ipairs(redis.call('smembers', tag .. '.tags')) do
		redis.call('del', key)
	end

	redis.call('del', tag .. '.tags')
end
`
