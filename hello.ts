class Greeter {
    greeting: string;
    constructor(message: string) {
        this.greeting = message;
    }
    greet() {
        return "Hello, " + this.greeting;
    }
}

var greeter = new Greeter("world");

var button = document.createElement('button');
button.textContent = "Say Hello";
button.onclick = function() {
    alert(greeter.greet());
}

document.body.appendChild(button);

apt-get install nodejs npm
npm install -g typescript 
ln -s /usr/bin/nodejs /usr/bin/node

/usr/bin/env: node: No such file or directory
nodejs /usr/local/lib/node_modules/typescript/bin/tsc /root/sources/go/hello.ts 
bash: node: command not found
nodejs /usr/local/lib/node_modules/typescript/bin/tsc /root/sources/go/hello.ts 

or
node built/local/tsc.js hello.ts


url(blob:https%3A//web.whatsapp.com/0bd61096-06c6-4b11-a3eb-0cb3335c4201)



// https://en.wikipedia.org/wiki/Comparison_of_e-book_formats
https://epubreader.codeplex.com/
http://stackoverflow.com/questions/21626219/convert-html-files-to-epub-files-programmatically-command-line-ubuntu
http://www.guidingtech.com/9661/difference-between-epub-mobi-azw-pdf-ebook-formats/


