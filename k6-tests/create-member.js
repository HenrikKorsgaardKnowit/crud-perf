import http from 'k6/http';
import { check, sleep } from 'k6';

// Lock request time out period exceeded
const virtualUsersCount = 50;

export const options = {
    vus: virtualUsersCount,
    iterations: 1000,
};

function generateShortHash() {
    //return 21;
    return Math.random().toString(36).substring(2, 15);
}

function getUserIndex() {
    return (__VU - 1) * (options.iterations / options.vus) + __ITER + 1;
}

export default function () {
    const userIndex = getUserIndex();
    const email = `T${generateShortHash()}user_${userIndex}@example.com`;
    const name = `T${generateShortHash()}User_${userIndex}`;

    const payload = {
        Email: email,
        Password: 'Test_User_1!',
        Name: name,
        PhoneNumber: '+4512345678',
        PostCode: '1000',
        City: 'Aarhus',
        DoesNotLiveInDenmark: 'false',
        SchoolId: '79710b5a-6f7e-4554-ad57-9abd310fa146',
        FieldOfStudyId: '9b272d25-f7aa-4333-80d6-9ec5eabf6b10',
        NoLongerAttendingTheCourse: 'false',
        TermsAccepted: 'true',
    };

    const res = http.post('http://localhost:3000/users', JSON.stringify(payload), {
        headers: {
            'Content-Type': '"application/json"',
        },
    });

    /*
    check(res, {
        'status is 302 (redirect)': (r) => r.status === 302,
    });*/ 

    check(res, {
        'status is 201 (created)': (r) => r.status === 201
    })

    console.log(`Created: ${email} | Status: ${res.status}`);
    sleep(0.1);
}