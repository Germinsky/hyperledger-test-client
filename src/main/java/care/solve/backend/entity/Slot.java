package care.solve.backend.entity;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Slot {

    private String slotId;
    private Long timeStart;
    private Long timeFinish;
    private Type availability;
    RegistrationInfo registrationInfo;

    enum Type {
        FREE,
        BUSY
    }

}
