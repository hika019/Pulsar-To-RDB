#bin/bash

if [ ! -e /pulsar/checked ]; then

    /pulsar/bin/pulsar-admin tenants create private
    /pulsar/bin/pulsar-admin namespaces create private/test-namespaces
    /pulsar/bin/pulsar-admin topics create persistent://private/test-namespaces/users

    touch /pulsar/checked
fi
