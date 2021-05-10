import com.dpb.netty.codec.protobuf.SubscribeReqProto;
import com.dpb.netty.codec.protobuf.SubscribeReqProto.SubscribeReq;
import com.google.protobuf.InvalidProtocolBufferException;
import java.util.ArrayList;
import java.util.List;

public class TestSubscribeReqProto {

  /**
   * 编码
   * @return
   */
  private static byte[] encode(SubscribeReqProto.SubscribeReq req){
    return req.toByteArray();
  }


  private static SubscribeReqProto.SubscribeReq decode(byte[] body)
      throws InvalidProtocolBufferException {
    return SubscribeReqProto.SubscribeReq.parseFrom(body);
  }

  /**
   * 构建SubscribeReq对象
   * @return
   */
  private static SubscribeReqProto.SubscribeReq createSubscribeReq(){
    SubscribeReqProto.SubscribeReq.Builder builder = SubscribeReqProto.SubscribeReq.newBuilder();
    builder.setSubReqID(1);
    builder.setUserName("bobo");
    builder.setPreductName("Netty");
    List<String> address = new ArrayList<>();
    address.add("beijing");
    address.add("guangzhou");
    address.add("shezheng");
    builder.addAllAddress(address);
    return builder.build();
  }

  public static void main(String[] args) throws InvalidProtocolBufferException {
    SubscribeReqProto.SubscribeReq req = createSubscribeReq();
    System.out.println("编码前:"+req.toString());
    SubscribeReq req2 = decode(encode(req));
    System.out.println("编码后:"+req);
    System.out.println("编码后:"+req2);
    System.out.println(req2.equals(req));
  }
}
