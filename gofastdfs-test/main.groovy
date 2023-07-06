import java.io.*
import okhttp3.*

String path = "http://172.22.90.1:8080/group1/upload";
OkHttpClient client = new OkHttpClient();
File file = new File("./fileserver.exe");

RequestBody fileBody = RequestBody.create(MediaType.parse("application/x-www-form-urlencoded;charset=utf-8") , file);
String name = "fileName";
try {
  name = URLEncoder.encode(name, "UTF-8");
} catch (UnsupportedEncodingException e1) {
  //TODO
}

MultipartBody mBody = new MultipartBody.Builder().setType(MultipartBody.FORM)
    .addFormDataPart("path" , "")
    .addFormDataPart("scene" , "")
    .addFormDataPart("output" , "json")
    .addFormDataPart("file" , name , fileBody)
    .build();


Request request = new Request.Builder().url(path).header("token","fjdsklfdef=").post(mBody).build();


try {
  Response response = client.newCall(request).execute();
  if(response.isSuccessful()){
    String result = response.body().string();
    System.out.println("result: \n "+result);
  }
  response.body().close();
} catch (Exception e) {
  System.out.println("异常！！！");
}
