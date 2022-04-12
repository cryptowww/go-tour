# Excel自动拆分工具

## 生成带logo的执行文件

1. 先准备favicon.ico
2. .rc文件，内容

> IDI_ICON1 ICON "favicon.ico"

3. 创建manifest文件, 命名：go文件名.exe.manifest

```xml

<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="x86"
    name="controls"
    type="win32"
></assemblyIdentity>
<dependency>
    <dependentAssembly>
        <assemblyIdentity
            type="win32"
            name="Microsoft.Windows.Common-Controls"
            version="6.0.0.0"
            processorArchitecture="*"
            publicKeyToken="6595b64144ccf1df"
            language="*"
        ></assemblyIdentity>
    </dependentAssembly>
</dependency>
</assembly>

```
4. 生成syso文件

> rsrc -manifest go文件名.exe.manifest -ico favicon.ico -o go文件名.syso


5. 编译

> go build

