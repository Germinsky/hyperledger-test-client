package care.solve.backend.transformer;

import care.solve.backend.entity.RegistrationInfo;
import care.solve.backend.entity.ScheduleProtos;
import org.springframework.stereotype.Service;

@Service
public class RegistrationInfoTransformer implements ProtoTransformer<RegistrationInfo, ScheduleProtos.RegistrationInfo> {

    @Override
    public ScheduleProtos.RegistrationInfo transformToProto(RegistrationInfo obj) {
        return ScheduleProtos.RegistrationInfo.newBuilder()
                .setDescription(obj.getDescription())
                .setPatientId(obj.getPatientId())
                .build();
    }

    @Override
    public RegistrationInfo transformFromProto(ScheduleProtos.RegistrationInfo proto) {
        return RegistrationInfo.builder()
                .description(proto.getDescription())
                .patientId(proto.getPatientId())
                .build();
    }
}