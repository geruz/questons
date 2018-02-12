resolve: {
        extensions: ['.js', '.jsx'],
        alias: config.alias,
        plugins: [{apply: function(resolver){
            resolver.plugin('new-resolve',
            function(shortRequest, callback) {
                const nameR = /^design\/([a-z]+):([0-9]\.[0-9]\.[0-9])?/
                const match = nameR.exec(shortRequest.request);
                if (match) {
                    const name = match[1];
                    const tag = match[2];
                    const compPath = path.join(path.resolve(), `./components/${name}`)
                    if (fs.existsSync(compPath)) {
                        shortRequest.request = `~/components/${name}`
                        callback()
                        return
                    } 
                    return
                }
                
            })
        }}],
