
import io.netty.channel.ChannelInboundHandlerAdapter;
import java.util.ArrayList;
import java.util.List;

import com.dpb.netty.codec.protobuf.SubscribeReqProto;

import io.netty.channel.ChannelHandlerAdapter;
import io.netty.channel.ChannelHandlerContext;

public class SubReqClientHandler extends ChannelInboundHandlerAdapter {

  /**
   * Creates a client-side handler.
   */
  public SubReqClientHandler() {
  }

  @Override
  public void channelActive(ChannelHandlerContext ctx) {
    for (int i = 0; i < 10; i++) {
      ctx.write(subReq(i));
    }
    ctx.flush();
  }

  private SubscribeReqProto.SubscribeReq subReq(int i) {
    SubscribeReqProto.SubscribeReq.Builder builder = SubscribeReqProto.SubscribeReq.newBuilder();
    builder.setSubReqID(i);
    builder.setUserName("bobo");
    builder.setPreductName("Netty Book For Protobuf");
    List<String> address = new ArrayList<>();
    address.add("NanJing");
    address.add("BeiJing");
    address.add("ShenZhen");
    builder.addAllAddress(address);
    return builder.build();
  }

  @Override
  public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
    System.out.println("Receive server response : [" + msg + "]");
  }

  @Override
  public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
    ctx.flush();
  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
    cause.printStackTrace();
    ctx.close();
  }
}