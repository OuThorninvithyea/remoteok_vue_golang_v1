import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useDropdownLocation() {
    const locationData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/dropdown/location')
            locationData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { locationData, loading, error }
}
