package care.solve.backend.controller;

import care.solve.backend.entity.PatientPrivate;
import care.solve.backend.service.PatientServiceWrapper;
import care.solve.backend.service.ScheduleServiceWrapper;
import care.solve.protocol.schedule.entity.PatientPublic;
import care.solve.protocol.schedule.entity.Schedule;
import com.google.protobuf.InvalidProtocolBufferException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.util.concurrent.ExecutionException;

@RestController
@RequestMapping("/patient")
public class PatientController {

    private PatientServiceWrapper patientServiceWrapper;
    private ScheduleServiceWrapper scheduleServiceWrapper;

    @Autowired
    public PatientController(PatientServiceWrapper patientServiceWrapper, ScheduleServiceWrapper scheduleServiceWrapper) {
        this.patientServiceWrapper = patientServiceWrapper;
        this.scheduleServiceWrapper = scheduleServiceWrapper;
    }

    @GetMapping("{patientId}")
    public PatientPublic get(@PathVariable String patientId) throws IOException {
        return patientServiceWrapper.get(patientId);
    }

    @PostMapping
    public PatientPublic create(@RequestBody PatientPrivate patient) throws InterruptedException, ExecutionException, InvalidProtocolBufferException {
        PatientPublic patientPublic = patientServiceWrapper.create(patient);
        Schedule patientsSchedule = Schedule.builder().ownerId(patientPublic.getId()).build();
        scheduleServiceWrapper.createSchedule(patientsSchedule);

        return patientPublic;
    }

}
